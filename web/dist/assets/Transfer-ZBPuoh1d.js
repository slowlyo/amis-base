import{aq as fe,s as ee,E as $,G as z,bS as ge,a3 as P,y as v,aK as ye,v as me,X as be,bb as Ce,aF as Oe,f as k,bT as Te,bJ as G,ay as Fe,aM as Se,O as Re,z as S,P as A,Q as r,aX as _e,bk as Ae,bU as J,bV as we,bW as D,bX as Q,bY as Y,bZ as Ie,ah as xe,bq as Me,bf as Z,aw as Pe,aA as Ee}from"./index-B1wf86tJ.js";var Ve=function(O){switch(typeof O){case"string":return fe(O,"options","inputValue","option");case"function":return O;default:return null}},ke=function(O){ee(o,O);function o(){return O!==null&&O.apply(this,arguments)||this}return o.prototype.reload=function(){var t=this.props.reloadOptions;t==null||t()},o.prototype.handleChange=function(t,n){return $(this,void 0,void 0,function(){var e,a,i,d,s,u,l,w,f,b,p,T,x,m,h,I,E,M,C;return z(this,function(g){switch(g.label){case 0:return e=this.props,a=e.onChange,i=e.joinValues,d=e.delimiter,s=e.valueField,u=e.extractValue,l=e.options,w=e.dispatchEvent,f=e.setOptions,b=e.selectMode,p=e.deferApi,T=e.deferField,x=T===void 0?"defer":T,m=t,h=l.concat(),Array.isArray(t)?(m=t.map(function(c){var F=J(l,D(c[s||"value"],s||"value"),{resolve:we(s),value:c[s]||"value"});if(!F)h.push(v(v({},c),{visible:!1}));else if(n){var V=Q(h,F);h=Y(h,F,1,v(v({},V),c))}return i||u?c[s||"value"]:c}),i&&(m=m.join(d||","))):t&&(m=i||u?t[s||"value"]:t,I=J(l,D(t[s||"value"],s||"value")),I?n&&(E=Q(h,I),h=Y(h,I,1,v(v({},E),t))):h.push(v(v({},t),{visible:!1}))),M=b==="tree"&&(!!p||!!Ie(l,function(c){return c.deferApi||c[x]})),(M===!0||h.length>l.length||n)&&f(h,!0),[4,w("change",xe(this.props,{value:m,options:l,items:l}))];case 1:return C=g.sent(),C!=null&&C.prevented?[2]:(a(m),[2])}})})},o.prototype.option2value=function(t){return t},o.prototype.handleSearch=function(t,n,e){return $(this,void 0,void 0,function(){var a,i,d,s,u,l,w,f,b,p,T,x,m,h,I,E,M;return z(this,function(C){switch(C.label){case 0:if(a=this.props,i=a.searchApi,d=a.options,s=a.labelField,u=a.valueField,l=a.env,w=a.data,f=a.translate,b=a.filterOption,!i)return[3,5];C.label=1;case 1:return C.trys.push([1,3,,4]),[4,l.fetcher(i,P(w,v({term:t},e||{})),{cancelExecutor:n})];case 2:if(p=C.sent(),!p.ok)throw new Error(f(p.msg||"networkError"));if(T=p.data.options||p.data.items||p.data,!Array.isArray(T))throw new Error(f("CRUD.invalidArray"));return x={},e&&(x={page:p.data.page,perPage:e.perPage,total:p.data.count}),[2,v({items:Pe(T,function(g){var c=null,F=g[u||"value"];return Array.isArray(d)&&F!==null&&F!==void 0&&(c=Ee(d,D(F,u)),g!=null&&g.children&&(c=v(v({},c),{children:g.children}))),c||g})},x)];case 3:return m=C.sent(),!l.isCancel(m)&&!i.silent&&l.notify("error",m.message),[2,{items:[]}];case 4:return[3,6];case 5:return t?(h=s||"label",I=u||"value",E={keys:[h,I]},b?(M=Ve(b),M?[2,{items:M(d,t,E)}]:(l.notify("error","自定义检索函数不符合要求"),[2,{items:[]}])):[2,{items:Me(d,function(g,c,F,V){return!!(Array.isArray(g.children)&&g.children.length||Z([g].concat(V),t,{keys:[s||"label",u||"value"],threshold:Z.rankings.CONTAINS}).length)},0,!0)}]):[2,{items:d}];case 6:return[2]}})})},o.prototype.handleResultSearch=function(t,n){var e=this.props,a=e.valueField,i=e.labelField,d=ge(t),s=n[i||"label"],u=n[a||"value"];return d.test(s)||d.test(u)},o.prototype.handlePageChange=function(t,n,e){var a=this.props,i=a.source,d=a.data,s=a.formItem,u=a.onChange,l=P(d,v({page:t??1,perPage:n??10},e?{pageDir:e}:{}));!s||!ye(s)||(me(i)?s.loadOptionsFromDataScope(i,l,u):be(i,l)&&s.loadOptions(i,l,void 0,!1,u,!1))},o.prototype.optionItemRender=function(t,n){var e=this.props,a=e.menuTpl,i=e.render,d=e.data;return i("item/".concat(n.index),a,{data:P(P(d,n),t)})},o.prototype.resultItemRender=function(t,n){var e=this.props,a=e.valueTpl,i=e.render,d=e.data;return i("value/".concat(n.index),a,{onChange:n.onChange,data:P(P(d,n),t)})},o.prototype.renderCell=function(t,n,e,a){var i=this.props,d=i.render,s=i.data,u=i.classnames,l=i.showInvalidMatch;return d("cell/".concat(e,"/").concat(a),v({type:"text",className:u({"is-invalid":l?n==null?void 0:n.__unmatched:!1})},t),{value:Ce(t.name,n),data:P(s,n)})},o.prototype.getRef=function(t){for(;t&&t.getWrappedInstance;)t=t.getWrappedInstance();this.tranferRef=t},o.prototype.onSelectAll=function(t){var n=this.props,e=n.dispatchEvent,a=n.data;e("selectAll",P(a,{items:t}))},o.prototype.doAction=function(t,n,e){var a,i,d,s,u,l=this.props,w=l.resetValue,f=l.onChange,b=l.formStore,p=l.store,T=l.name;switch(t.actionType){case"clear":f==null||f("");break;case"reset":f==null||f((d=(i=Oe((a=b==null?void 0:b.pristine)!==null&&a!==void 0?a:p==null?void 0:p.pristine,T))!==null&&i!==void 0?i:w)!==null&&d!==void 0?d:"");break;case"selectAll":(s=this.tranferRef)===null||s===void 0||s.selectAll();break;case"clearSearch":{(u=this.tranferRef)===null||u===void 0||u.clearSearch(n);break}}},o.prototype.render=function(){var t,n,e=this.props,a=e.className;e.style;var i=e.classnames,d=e.selectedOptions,s=e.showArrow,u=e.sortable,l=e.selectMode,w=e.columns,f=e.loading,b=e.searchable,p=e.searchResultMode,T=e.searchResultColumns,x=e.deferLoad,m=e.leftMode,h=e.rightMode,I=e.disabled,E=e.selectTitle,M=e.resultTitle,C=e.menuTpl,g=e.valueTpl,c=e.searchPlaceholder,F=e.resultListModeFollowSelect,V=F===void 0?!1:F,te=e.resultSearchPlaceholder,L=e.resultSearchable,ne=L===void 0?!1:L,ae=e.statistics,re=e.labelField,le=e.valueField,ie=e.virtualThreshold,U=e.itemHeight,H=e.loadingConfig,oe=e.showInvalidMatch,se=e.onlyChildren,de=e.mobileUI,ue=e.noResultsText,R=e.pagination,y=e.formItem,N=e.env,j=e.popOverContainer,pe=e.data,K=e.autoCheckChildren,ce=K===void 0?!0:K,W=e.initiallyOpen,he=W===void 0?!0:W,ve=e.testIdBuilder,B=this.props,_=B.options,X=B.leftOptions,q=B.leftDefaultValue;return l==="associated"&&_&&_.length&&_[0].leftOptions&&Array.isArray(_[0].children)&&(X=_[0].leftOptions,q=(t=_[0].leftDefaultValue)!==null&&t!==void 0?t:q,_=_[0].children),k.createElement("div",{className:i("TransferControl",a)},k.createElement(Te,{onlyChildren:se,value:d,options:_,accumulatedOptions:(n=y==null?void 0:y.accumulatedOptions)!==null&&n!==void 0?n:[],disabled:I,onChange:this.handleChange,option2value:this.option2value,sortable:u,showArrow:s,selectMode:l,searchResultMode:p,searchResultColumns:T,columns:w,onSearch:b?this.handleSearch:void 0,onDeferLoad:x,leftOptions:X,leftMode:m,rightMode:h,cellRender:this.renderCell,selectTitle:E,resultTitle:M,resultListModeFollowSelect:V,onResultSearch:this.handleResultSearch,searchPlaceholder:c,resultSearchable:ne,resultSearchPlaceholder:te,statistics:ae,labelField:re,valueField:le,optionItemRender:C?this.optionItemRender:void 0,resultItemRender:g?this.resultItemRender:void 0,onSelectAll:this.onSelectAll,onRef:this.getRef,virtualThreshold:ie,itemHeight:G(U)>0?G(U):void 0,loadingConfig:H,showInvalidMatch:oe,mobileUI:de,noResultsText:ue,pagination:v(v({},Fe(R,["className","layout","perPageAvailable","popOverContainerSelector"])),{enable:(R&&R.enable!==void 0?!!(typeof R.enable=="string"?Se(R.enable,pe):R.enable):!!(y!=null&&y.enableSourcePagination))&&(!l||l==="list"||l==="table")&&_.length>0,maxButtons:Number.isInteger(R==null?void 0:R.maxButtons)?R.maxButtons:5,page:y==null?void 0:y.sourcePageNum,perPage:y==null?void 0:y.sourcePerPageNum,total:y==null?void 0:y.sourceTotalNum,popOverContainer:j??(N==null?void 0:N.getModalContainer)}),onPageChange:this.handlePageChange,initiallyOpen:he,autoCheckChildren:ce,testIdBuilder:ve}),k.createElement(Re,{overlay:!0,key:"info",loadingConfig:H,show:f}))},o.defaultProps={multiple:!0},S([A,r("design:type",Function),r("design:paramtypes",[Object,Boolean]),r("design:returntype",Promise)],o.prototype,"handleChange",null),S([A,r("design:type",Function),r("design:paramtypes",[Object]),r("design:returntype",void 0)],o.prototype,"option2value",null),S([A,r("design:type",Function),r("design:paramtypes",[String,Function,Object]),r("design:returntype",Promise)],o.prototype,"handleSearch",null),S([A,r("design:type",Function),r("design:paramtypes",[String,Object]),r("design:returntype",void 0)],o.prototype,"handleResultSearch",null),S([A,r("design:type",Function),r("design:paramtypes",[Number,Number,String]),r("design:returntype",void 0)],o.prototype,"handlePageChange",null),S([A,r("design:type",Function),r("design:paramtypes",[Object,Object]),r("design:returntype",void 0)],o.prototype,"optionItemRender",null),S([A,r("design:type",Function),r("design:paramtypes",[Object,Object]),r("design:returntype",void 0)],o.prototype,"resultItemRender",null),S([A,r("design:type",Function),r("design:paramtypes",[Object,Object,Number,Number]),r("design:returntype",void 0)],o.prototype,"renderCell",null),S([A,r("design:type",Function),r("design:paramtypes",[Object]),r("design:returntype",void 0)],o.prototype,"getRef",null),S([A,r("design:type",Function),r("design:paramtypes",[Array]),r("design:returntype",void 0)],o.prototype,"onSelectAll",null),S([_e(),r("design:type",Function),r("design:paramtypes",[]),r("design:returntype",void 0)],o.prototype,"render",null),o}(k.Component),Ne=function(O){ee(o,O);function o(){return O!==null&&O.apply(this,arguments)||this}return o}(ke),Be=Ae({type:"transfer"})(Ne);export{ke as BaseTransferRenderer,Ne as TransferRender,Be as default,Ve as getCustomFilterOption};