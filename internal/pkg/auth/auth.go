package auth

import (
	"amis-base/internal/models"
	"amis-base/internal/pkg/db"
	"amis-base/internal/pkg/helper"
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

	db.GetDB().Create(&models.Token{
		TableName:  tableName,
		UserId:     userId,
		Token:      helper.Sha256Hash(token),
		LastUsedAt: time.Now(),
	})

	return token
}

// CleanTokenByUserId 清除用户 token
func CleanTokenByUserId(userId uint) {
	db.GetDB().Where("user_id = ?", userId).Delete(&models.Token{})
}

// CleanExpiredToken 清除过期的 token
func CleanExpiredToken() {
	expireTime := viper.GetInt("admin.auth.token_expire")

	if expireTime == 0 {
		return
	}

	latestTime := time.Now().Add(-time.Duration(expireTime) * time.Second)

	db.GetDB().Where("last_used_at < ?", latestTime).Delete(&models.Token{})
}

// Hash 加密密码
func Hash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return ""
	}

	return string(hashedPassword)
}

// CheckHash 校验密码
func CheckHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
