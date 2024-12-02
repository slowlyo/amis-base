package auth

import (
	"amis-base/internal/models"
	"amis-base/internal/pkg/cache"
	"amis-base/internal/pkg/db"
	"amis-base/internal/pkg/helper"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// GenerateToken 生成 token
func GenerateToken(tableName string, userId uint) string {
	token := uuid.NewString()

	// 单点登录, 清除其余 token
	if viper.GetBool("admin.auth.single_sign_on") {
		go CleanTokenByUserId(userId)
	}

	db.Query().Create(&models.Token{
		TableName:  tableName,
		UserId:     userId,
		Token:      helper.Sha256Hash(token),
		LastUsedAt: time.Now(),
	})

	return token
}

// CleanTokenByUserId 清除用户 token
func CleanTokenByUserId(userId uint) {
	db.Query().Where("user_id = ?", userId).Delete(&models.Token{})
}

// CleanExpiredToken 清除过期的 token
func CleanExpiredToken() {
	expireTime := viper.GetInt("admin.auth.token_expire")

	if expireTime == 0 {
		return
	}

	latestTime := time.Now().Add(-time.Duration(expireTime) * time.Second)

	db.Query().Where("last_used_at < ?", latestTime).Delete(&models.Token{})
}

// Hash 加密密码
func Hash(password string) string {
	return helper.BcryptString(password)
}

// CheckHash 校验密码
func CheckHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// QueryToken 查询 token 信息
func QueryToken(tableName, token string) *models.Token {
	// 缓存 key
	cacheKey := fmt.Sprintf("tokens:%s:%s", tableName, helper.Sha256Hash(token))

	// 更新最后使用时间
	updateLastUsedAt := func(tokenModel models.Token) {
		db.Query().Model(&tokenModel).Update("last_used_at", time.Now())
	}

	// 缓存命中
	if cachedToken := cache.GetObject[models.Token](cacheKey); cachedToken.ID > 0 {
		go updateLastUsedAt(cachedToken)

		return &cachedToken
	}

	// token 过期时间
	tokenExpire := time.Duration(viper.GetInt("admin.auth.token_expire")) * time.Second

	var tokenModel models.Token

	result := db.Query().
		Where("table_name = ?", tableName).
		Where("token = ?", helper.Sha256Hash(token)).
		Where("last_used_at > ?", time.Now().Add(-tokenExpire)).
		First(&tokenModel)

	if result.RowsAffected == 0 {
		return nil
	}

	// 缓存 token
	_ = cache.SetObject(cacheKey, tokenModel, tokenExpire)

	go updateLastUsedAt(tokenModel)

	return &tokenModel
}

// RemoveToken 删除 token
func RemoveToken(token string) {
	db.Query().Where("token = ?", helper.Sha256Hash(token)).Delete(&models.Token{})
}
