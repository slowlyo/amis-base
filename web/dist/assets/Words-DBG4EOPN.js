import{s as g,h as c,F as h,cx as b,j as m,dW as _,y as p,w as P,z as y,P as W,Q as f,f as B,D as x,c5 as E}from"./index-kH63o2-N.js";function F(l,t,e){var r=e.type,s=e.labelField,a=s===void 0?"label":s,n=e.options,o=n===void 0?[]:n,i=e.enableNodePath,d=e.hideNodePathLabel,u=e.pathSeparator,T=u===void 0?"/":u;if(i||r==="nested-select"&&!d){var v=E(o,l,!0);return"".concat(v?v.map(function(N){return"".concat(N[a||"label"])}).join(" ".concat(T," ")):l[a||"label"])}return b(l[a])||"选项".concat(t)}var A=function(l){g(t,l);function t(){var e=l!==null&&l.apply(this,arguments)||this;return e.state={isExpend:!1},e}return t.prototype.toggleExpend=function(){this.setState({isExpend:!this.state.isExpend})},t.prototype.getLimit=function(e){var r=this.props.limit;return r??(Array.isArray(e)?10:200)},t.prototype.renderContent=function(e){var r=this.props,s=r.delimiter,a=r.inTag,n=r.classnames;if(!Array.isArray(e))return e;if(!a){var o=e.length-1;return e.map(function(i,d){return c(h,{children:[b(i),d===o?"":s||"， "]})})}return e.map(function(i,d){return m(_,{...p({key:d,label:i,className:"mb-1"},typeof a=="object"?p(p({},a),{className:n(a.className)}):void 0)})})},t.prototype.renderAll=function(e,r){r===void 0&&(r=!1);var s=this.props,a=s.collapseButtonText,n=a===void 0?"收起":a,o=s.collapseButton,i=s.render;return c(h,{children:[this.renderContent(e),r?i("collapseBtn",{type:"button",level:"link",className:"ml-1 v-baseline"},p(p({onClick:this.toggleExpend},o),{label:n})):null]})},t.prototype.renderPart=function(e){var r=this.props,s=r.expendButtonText,a=s===void 0?"展开":s,n=r.expendButton,o=r.render,i=this.getLimit(e),d=Array.isArray(e)?e.slice(0,i):e.toString().slice(0,i);return c(h,{children:[this.renderContent(d)," ...",o("collapseBtn",{type:"button",level:"link",className:"ml-1 v-baseline"},p(p({onClick:this.toggleExpend},n),{label:a}))]})},t.prototype.getWords=function(){var e=this,r=this.props,s=r.selectedOptions,a=s===void 0?[]:s,n=r.words,o=r.data,i;return typeof n=="string"&&(i=P(n,o,"| raw")),i||((a==null?void 0:a.length)>0?a.map(function(d,u){return F(d,u,e.props)}):null)},t.prototype.render=function(){var e=this.props,r=e.classnames,s=e.className,a=e.style,n=this.getWords();if(!n)return null;var o=this.getLimit(n),i;return!o||Array.isArray(n)&&n.length<=o||!Array.isArray(n)&&n.toString().length<=o?i=this.renderAll(n):i=this.state.isExpend?this.renderAll(n,!0):this.renderPart(n),m("div",{className:r("Words-field",s),style:a,children:i})},t.defaultProps={inTag:!1},y([W,f("design:type",Function),f("design:paramtypes",[]),f("design:returntype",void 0)],t.prototype,"toggleExpend",null),t}(B.Component),L=function(l){g(t,l);function t(){return l!==null&&l.apply(this,arguments)||this}return t=y([x({type:"words"})],t),t}(A),R=function(l){g(t,l);function t(){return l!==null&&l.apply(this,arguments)||this}return t.defaultProps={inTag:!0},t=y([x({type:"tags"})],t),t}(A);export{R as TagsRenderer,A as WordsField,L as WordsRenderer};