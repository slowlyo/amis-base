import{s,ai as d,f as r,z as m,D as p}from"./index-Cuf9I9Av.js";var f=function(l){s(e,l);function e(){return l!==null&&l.apply(this,arguments)||this}return e.prototype.render=function(){var o=this.props,n=o.className,i=o.style,a=o.classnames,u=o.defaultColor,c=o.showValue,t=d(this.props)||u;return r.createElement("div",{className:a("ColorField",n),style:i},r.createElement("i",{className:a("ColorField-previewIcon"),style:{backgroundColor:t}}),c&&t?r.createElement("span",{className:a("ColorField-value")},t):null)},e.defaultProps={className:"",defaultColor:"",showValue:!0},e}(r.Component),h=function(l){s(e,l);function e(){return l!==null&&l.apply(this,arguments)||this}return e=m([p({type:"color"})],e),e}(f);export{f as ColorField,h as ColorFieldRenderer};