import{s as V,ah as B,a3 as C,y as c,cQ as D,aF as Q,f,cR as $,bJ as R,O as J,z as u,P as F,Q as s,aX as X,bk as q}from"./index-B1wf86tJ.js";import{BaseTabsTransferRenderer as G}from"./TabsTransfer-DNFAT4L-.js";import"./Transfer-ZBPuoh1d.js";var ee=function(v){V(o,v);function o(){var t=v!==null&&v.apply(this,arguments)||this;return t.state={activeKey:0},t}return o.prototype.dispatchEvent=function(t){var e=this.props,a=e.dispatchEvent,r=e.value;a(t,B(this.props,{value:r}))},o.prototype.optionItemRender=function(t,e){var a=this.props,r=a.menuTpl,i=a.render,d=a.data,n=a.classnames,l=arguments[2]||{};return r?i("item/".concat(e.index),r,{data:C(C(d,c(c({},e),l)),t)}):D.itemRender(t,c(c({},e),{classnames:n}))},o.prototype.doAction=function(t){var e,a,r,i=this.props,d=i.resetValue,n=i.onChange,l=i.formStore,p=i.store,h=i.name;switch(t.actionType){case"clear":n==null||n("");break;case"reset":n==null||n((r=(a=Q((e=l==null?void 0:l.pristine)!==null&&e!==void 0?e:p==null?void 0:p.pristine,h))!==null&&a!==void 0?a:d)!==null&&r!==void 0?r:"");break}},o.prototype.render=function(){var t=this,e=this.props,a=e.className;e.style;var r=e.classnames,i=e.options,d=e.selectedOptions,n=e.sortable,l=e.loading,p=e.searchResultMode,h=e.showArrow,O=e.deferLoad,w=e.disabled,k=e.selectTitle,E=e.resultTitle,S=e.pickerSize,I=e.leftMode,x=e.leftOptions,g=e.itemHeight,M=e.virtualThreshold,P=e.loadingConfig,b=e.labelField,z=b===void 0?"label":b,T=e.valueField,A=T===void 0?"value":T,y=e.deferField,N=y===void 0?"defer":y,j=e.mobileUI,m=e.env,H=e.maxTagCount,K=e.overflowTagPopover,L=e.placeholder,_=e.initiallyOpen,U=_===void 0?!0:_;return f.createElement("div",{className:r("TabsTransferControl",a)},f.createElement($,{activeKey:this.state.activeKey,onTabChange:this.onTabChange,placeholder:L,value:d,disabled:w,options:i,onChange:this.handleChange,option2value:this.option2value,sortable:n,searchResultMode:p,onSearch:this.handleTabSearch,showArrow:h,onDeferLoad:O,selectTitle:k,resultTitle:E,size:S,leftMode:I,leftOptions:x,optionItemRender:this.optionItemRender,resultItemRender:this.resultItemRender,onFocus:function(){return t.dispatchEvent("focus")},onBlur:function(){return t.dispatchEvent("blur")},itemHeight:R(g)>0?R(g):void 0,virtualThreshold:M,labelField:z,valueField:A,deferField:N,mobileUI:j,popOverContainer:m==null?void 0:m.getModalContainer,maxTagCount:H,overflowTagPopover:K,initiallyOpen:U}),f.createElement(J,{loadingConfig:P,overlay:!0,key:"info",show:l}))},u([F,s("design:type",Function),s("design:paramtypes",[String]),s("design:returntype",void 0)],o.prototype,"dispatchEvent",null),u([F,s("design:type",Function),s("design:paramtypes",[Object,Object]),s("design:returntype",void 0)],o.prototype,"optionItemRender",null),u([X(),s("design:type",Function),s("design:paramtypes",[]),s("design:returntype",void 0)],o.prototype,"render",null),o=u([q({type:"tabs-transfer-picker"})],o),o}(G);export{ee as TabsTransferPickerRenderer};