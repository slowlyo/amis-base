import{s as a,ca as o,f as u,z as i,a1 as l}from"./index-Cuf9I9Av.js";var s=function(n){a(t,n);function t(e){var r=n.call(this,e)||this;return e.value||r.setValue(),r}return t.prototype.componentDidUpdate=function(e){!e.value&&e.formInited!==!1&&this.setValue()},t.prototype.setValue=function(){var e=this.props,r=o();e.length&&(r=r.substring(0,e.length)),e.onChange(r)},t.prototype.render=function(){return null},t}(u.Component),d=function(n){a(t,n);function t(){return n!==null&&n.apply(this,arguments)||this}return t=i([l({type:"uuid",wrap:!1,sizeMutable:!1})],t),t}(s);export{d as UUIDControlRenderer,s as default};