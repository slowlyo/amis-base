import{s as c,K as _,ai as C,f as p,dh as L,z as h,P as T,Q as o,D as R,$ as F}from"./index-kH63o2-N.js";var I=function(a){c(t,a);function t(){return a!==null&&a.apply(this,arguments)||this}return t.prototype.handleClick=function(e){var n=this.props,r=n.env,i=n.href,l=n.blank,s=n.body;r==null||r.tracker({eventType:"url",eventData:{url:i,blank:l,label:s}},this.props)},t.prototype.getHref=function(){},t.prototype.render=function(){var e=this.props,n=e.className,r=e.style,i=e.body,l=e.href;e.classnames;var s=e.blank,u=e.disabled,f=e.htmlTarget,y=e.data,b=e.render,k=e.translate,m=e.title,g=e.icon,v=e.rightIcon,d=(typeof l=="string"&&l?_(l,y,"| raw"):void 0)||C(this.props);return p.createElement(L,{className:n,style:r,href:d,disabled:u,title:m,htmlTarget:f||(s?"_blank":"_self"),icon:g,rightIcon:v,onClick:this.handleClick},i?b("body",i):d||k("link"))},t.defaultProps={blank:!0,disabled:!1,htmlTarget:""},h([T,o("design:type",Function),o("design:paramtypes",[Object]),o("design:returntype",void 0)],t.prototype,"handleClick",null),t}(p.Component),P=function(a){c(t,a);function t(){return a!==null&&a.apply(this,arguments)||this}return t=h([R({type:"link"}),F],t),t}(I);export{I as LinkCmpt,P as LinkFieldRenderer};