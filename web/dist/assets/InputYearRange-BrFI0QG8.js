import{s as D,t as R,f as v,ag as _,c1 as Y,y as I,aV as g,c2 as y,z as h,aZ as P,Q as m,a1 as b}from"./index-Cuf9I9Av.js";import E from"./InputDateRange-ioPjeNGj.js";var N=function(t){D(e,t);function e(){return t!==null&&t.apply(this,arguments)||this}return e.prototype.render=function(){var a=this.props,F=a.className;a.style;var l=a.classPrefix,u=a.minDate,d=a.maxDate,p=a.minDuration,c=a.maxDuration,r=a.data,o=a.format,f=a.mobileUI,i=a.valueFormat,x=a.inputFormat,C=a.displayFormat,n=a.env,s=R(a,["className","style","classPrefix","minDate","maxDate","minDuration","maxDuration","data","format","mobileUI","valueFormat","inputFormat","displayFormat","env"]);return v.createElement("div",{className:_("".concat(l,"DateRangeControl"),F)},v.createElement(Y,I({viewMode:"years",mobileUI:f,valueFormat:i||o,displayFormat:C||x,classPrefix:l,popOverContainer:f?n==null?void 0:n.getModalContainer:s.popOverContainer||n.getModalContainer,popOverContainerSelector:s.popOverContainerSelector,onRef:this.getRef,data:r},s,{minDate:u?g(u,r,i||o):void 0,maxDate:d?g(d,r,i||o):void 0,minDuration:p?y(p):void 0,maxDuration:c?y(c):void 0,onChange:this.handleChange,onFocus:this.dispatchEvent,onBlur:this.dispatchEvent})))},h([P(),m("design:type",Function),m("design:paramtypes",[]),m("design:returntype",void 0)],e.prototype,"render",null),e}(E),M=function(t){D(e,t);function e(){return t!==null&&t.apply(this,arguments)||this}return e.defaultProps={format:"X",inputFormat:"YYYY",joinValues:!0,delimiter:",",ranges:"thisyear,prevyear",shortcuts:"thisyear,prevyear",animation:!0},e=h([b({type:"input-year-range"})],e),e}(N);export{M as YearRangeControlRenderer,N as default};