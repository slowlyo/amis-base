import{s as k,E as N,G as A,f as r,H as C,J as x,K as w,L as F,M as D,N as H,O,z as B,P,Q as m,T as L,D as j,U as R,V as T,W as S,X as U,Y as W}from"./index-kH63o2-N.js";var I=function(g){k(l,g);function l(e){var t=this,n,o,a;t=g.call(this,e)||this;var i=e.store;return i.syncProps(e,void 0,["pages"]),i.updateActivePage(Object.assign({},(n=e.env)!==null&&n!==void 0?n:{},{showFullBreadcrumbPath:(o=e.showFullBreadcrumbPath)!==null&&o!==void 0?o:!1,showBreadcrumbHomePath:(a=e.showBreadcrumbHomePath)!==null&&a!==void 0?a:!0})),e.env.watchRouteChange&&(t.unWatchRouteChange=e.env.watchRouteChange(function(){var c,u,d;return i.updateActivePage(Object.assign({},(c=e.env)!==null&&c!==void 0?c:{},{showFullBreadcrumbPath:(u=e.showFullBreadcrumbPath)!==null&&u!==void 0?u:!1,showBreadcrumbHomePath:(d=e.showBreadcrumbHomePath)!==null&&d!==void 0?d:!0}))})),t}return l.prototype.componentDidMount=function(){return N(this,void 0,void 0,function(){var e,t,n,o;return A(this,function(a){switch(a.label){case 0:return e=this.props,t=e.data,n=e.dispatchEvent,[4,n("init",t,this)];case 1:return o=a.sent(),o!=null&&o.prevented?[2]:(this.reload(),[2])}})})},l.prototype.componentDidUpdate=function(e){var t,n,o;return N(this,void 0,void 0,function(){var a,i;return A(this,function(c){return a=this.props,i=a.store,i.syncProps(a,e,["pages"]),T(e.api,a.api,e.data,a.data)?this.reload():a.location&&a.location!==e.location&&i.updateActivePage(Object.assign({},(t=a.env)!==null&&t!==void 0?t:{},{showFullBreadcrumbPath:(n=a.showFullBreadcrumbPath)!==null&&n!==void 0?n:!1,showBreadcrumbHomePath:(o=a.showBreadcrumbHomePath)!==null&&o!==void 0?o:!0})),[2]})})},l.prototype.componentWillUnmount=function(){var e;(e=this.unWatchRouteChange)===null||e===void 0||e.call(this)},l.prototype.reload=function(e,t,n,o,a){return N(this,void 0,void 0,function(){var i,c,u,d,v,s,b,f,E,p;return A(this,function(h){switch(h.label){case 0:return t?[2,this.receive(t,void 0,a)]:(i=this.props,c=i.api,u=i.store,d=i.env,v=i.showFullBreadcrumbPath,s=v===void 0?!1:v,b=i.showBreadcrumbHomePath,f=b===void 0?!0:b,E=i.locale,U(c,u.data)?[4,u.fetchInitData(c,u.data,{})]:[3,2]);case 1:p=h.sent(),d.replaceText&&(p.data=W(p.data,d.replaceText,d.replaceTextIgnoreKeys)),p!=null&&p.data.pages&&(p.data=S(p.data,E),u.setPages(p.data.pages),u.updateActivePage(Object.assign({},d??{},{showFullBreadcrumbPath:s,showBreadcrumbHomePath:f}))),h.label=2;case 2:return[2,u.data]}})})},l.prototype.receive=function(e,t,n){return N(this,void 0,void 0,function(){var o;return A(this,function(a){return o=this.props.store,o.updateData(e,void 0,n),[2,this.reload()]})})},l.prototype.resolveDefinitions=function(e){var t,n=this.props,o=n.resolveDefinitions,a=n.store,i=(t=a.schema)===null||t===void 0?void 0:t.definitions;return(i==null?void 0:i[e])||o(e)},l.prototype.handleNavClick=function(e){e.preventDefault();var t=this.props.env,n=e.currentTarget.getAttribute("href");t.jumpTo(n,void 0,this.props.data)},l.prototype.renderHeader=function(){var e=this.props,t=e.classnames,n=e.brandName,o=e.header,a=e.render,i=e.store,c=e.logo,u=e.env;return!o&&!c&&!n?null:r.createElement(r.Fragment,null,r.createElement("div",{className:t("Layout-brandBar")},r.createElement("div",{onClick:i.toggleOffScreen,className:t("Layout-offScreenBtn")},r.createElement("i",{className:"bui-icon iconfont icon-collapse"})),r.createElement("div",{className:t("Layout-brand")},c&&~c.indexOf("<svg")?r.createElement(C,{className:t("AppLogo-html"),html:c,filterHtml:u.filterHtml}):c?r.createElement("img",{className:t("AppLogo"),src:c}):r.createElement("span",{className:"visible-folded "},n==null?void 0:n.substring(0,1)),r.createElement("span",{className:"hidden-folded m-l-sm"},n))),r.createElement("div",{className:t("Layout-headerBar")},r.createElement("a",{onClick:i.toggleFolded,type:"button",className:t("AppFoldBtn")},r.createElement("i",{className:"fa fa-".concat(i.folded?"indent":"dedent"," fa-fw")})),o?a("header",o):null))},l.prototype.renderAside=function(){var e=this,t=this.props,n=t.store,o=t.env,a=t.asideBefore,i=t.asideAfter,c=t.render,u=t.data;return r.createElement(r.Fragment,null,a?c("aside-before",a):null,r.createElement(x,{navigations:n.navigations,renderLink:function(d,v){var s=d.link;d.active;var b=d.toggleExpand,f=d.classnames,E=d.depth,p=d.subHeader,h=[];if(s.visible===!1)return null;!p&&s.children&&s.children.some(function(y){return y==null?void 0:y.visible})&&h.push(r.createElement("span",{key:"expand-toggle",className:f("AsideNav-itemArrow"),onClick:function(y){return b(s,y)}}));var _=typeof s.badge=="string"?w(s.badge,u):s.badge;return _!=null&&h.push(r.createElement("b",{key:"badge",className:f("AsideNav-itemBadge",s.badgeClassName||"bg-info")},_)),!p&&s.icon?h.push(r.createElement(F,{key:"icon",cx:f,icon:s.icon,className:"AsideNav-itemIcon"})):n.folded&&E===1&&!p&&h.push(r.createElement("i",{key:"icon",className:f("AsideNav-itemIcon",s.children?"fa fa-folder":"fa fa-info")})),h.push(r.createElement("span",{className:f("AsideNav-itemLabel"),key:"label"},typeof s.label=="string"?w(s.label,u):s.label)),s.path?/^https?\:/.test(s.path)?r.createElement("a",{target:"_blank",key:"link",href:s.path,rel:"noopener"},h):r.createElement("a",{key:"link",onClick:e.handleNavClick,href:s.path||s.children&&s.children[0].path},h):r.createElement("a",{key:"link",onClick:s.children?function(){return b(s)}:void 0},h)},isActive:function(d){return!!o.isCurrentUrl(d==null?void 0:d.path,d)}}),i?c("aside-before",i):null)},l.prototype.renderFooter=function(){var e=this.props,t=e.render,n=e.footer;return n?t("footer",n):null},l.prototype.render=function(){var e=this,t,n=this.props,o=n.classnames,a=n.store,i=n.render,c=n.showBreadcrumb,u=c===void 0?!0:c,d=n.loadingConfig;return r.createElement(D,{header:this.renderHeader(),aside:this.renderAside(),footer:this.renderFooter(),folded:a.folded,offScreen:a.offScreen,contentClassName:o("AppContent")},a.activePage&&a.schema?r.createElement(r.Fragment,null,u&&a.bcn.length?r.createElement("ul",{className:o("AppBcn")},a.bcn.map(function(v,s){return r.createElement("li",{key:s,className:o("AppBcn-item")},v.path?r.createElement("a",{href:v.path,onClick:e.handleNavClick},v.label):s!==a.bcn.length-1?r.createElement("a",null,v.label):v.label)})):null,r.createElement("div",{className:o("AppBody")},i("page",a.schema,{key:"".concat((t=a.activePage)===null||t===void 0?void 0:t.id,"-").concat(a.schemaKey),data:a.pageData,resolveDefinitions:this.resolveDefinitions}))):a.pages&&!a.activePage?r.createElement(H,null,r.createElement("div",{className:"text-center"},"页面不存在")):null,r.createElement(O,{loadingConfig:d,overlay:!0,show:a.loading||!a.pages,size:"lg"}))},l.propsList=["brandName","logo","header","asideBefore","asideAfter","pages","footer"],l.defaultProps={},B([P,m("design:type",Function),m("design:paramtypes",[String]),m("design:returntype",void 0)],l.prototype,"resolveDefinitions",null),B([P,m("design:type",Function),m("design:paramtypes",[Object]),m("design:returntype",void 0)],l.prototype,"handleNavClick",null),l}(r.Component),K=function(g){k(l,g);function l(e,t){var n=g.call(this,e)||this,o=t;return o.registerComponent(n),n}return l.prototype.componentWillUnmount=function(){var e=this.context;e.unRegisterComponent(this),g.prototype.componentWillUnmount.call(this)},l.prototype.setData=function(e,t){return this.props.store.updateData(e,void 0,t)},l.prototype.getData=function(){var e=this.props.store;return e.data},l.contextType=L,l=B([j({type:"app",storeType:R.name}),m("design:paramtypes",[Object,Object])],l),l}(I);export{I as App,K as default};