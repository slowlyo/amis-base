const __vite__mapDeps=(i,m=__vite__mapDeps,d=(m.f||(m.f=["assets/Markdown-BqoNLez6.js","assets/index-Cuf9I9Av.js","assets/index-DHDVZwRb.css"])))=>i.map(i=>d[i]);
import{s as h,ai as l,v as u,w as f,V as m,E as v,G as _,f as p,ds as w,z as y,D as C,o as k,X as E}from"./index-Cuf9I9Av.js";function g(){return k(()=>import("./Markdown-BqoNLez6.js"),__vite__mapDeps([0,1,2])).then(function(o){return o.default})}var M=function(o){h(n,o);function n(e){var t=o.call(this,e)||this,a=t.props,r=a.name,i=a.data,s=a.src;if(s)t.state={content:""},t.updateContent();else{var c=l(t.props)||(r&&u(r)?f(r,i,"| raw"):null);t.state={content:c}}return t}return n.prototype.componentDidUpdate=function(e){var t=this.props;t.src?m(e.src,t.src,e.data,t.data)&&this.updateContent():this.updateContent()},n.prototype.updateContent=function(){return v(this,void 0,void 0,function(){var e,t,a,r,i,s,c;return _(this,function(d){switch(d.label){case 0:return e=this.props,t=e.name,a=e.data,r=e.src,i=e.env,r&&E(r,a)?[4,i.fetcher(r,a)]:[3,2];case 1:return s=d.sent(),typeof s=="string"?this.setState({content:s}):typeof s=="object"&&s.data?this.setState({content:s.data}):console.error("markdown response error",s),[3,3];case 2:c=l(this.props)||(t&&u(t)?f(t,a,"| raw"):null),c!==this.state.content&&this.setState({content:c}),d.label=3;case 3:return[2]}})})},n.prototype.render=function(){var e=this.props,t=e.className,a=e.style,r=e.classnames,i=e.options;return p.createElement("div",{className:r("Markdown",t),style:a},p.createElement(w,{getComponent:g,content:this.state.content||"",options:i}))},n}(p.Component),V=function(o){h(n,o);function n(){return o!==null&&o.apply(this,arguments)||this}return n=y([C({type:"markdown"})],n),n}(M);export{M as Markdown,V as MarkdownRenderer};