import{s as k,ai as T,dH as N,f as i,L as d,aQ as E,z as r,P as u,Q as n,D as S,K as f}from"./index-Cuf9I9Av.js";var F=function(p){k(a,p);function a(){var t=p!==null&&p.apply(this,arguments)||this;return t.state={src:T(t.props,function(e){return e.src?f(e.src,e.data,"| raw"):void 0})||"",isReady:!1,muted:!1,playing:!1,played:0,seeking:!1,volume:.8,prevVolume:.8,loaded:0,playbackRate:1,showHandlePlaybackRate:!1,showHandleVolume:!1},t}return a.prototype.componentWillUnmount=function(){clearTimeout(this.progressTimeout),clearTimeout(this.durationTimeout)},a.prototype.componentDidMount=function(){var t=this.props.autoPlay,e=!!t;this.setState({playing:e},this.progress)},a.prototype.componentDidUpdate=function(t){var e=this,s=this.props;N(s,t,function(o){return e.setState({src:o,playing:!1},function(){e.audio.load(),e.progress()})},function(o){return o.src?f(o.src,o.data,"| raw"):void 0})},a.prototype.progress=function(){if(clearTimeout(this.progressTimeout),this.state.src&&this.audio){var t=this.audio.currentTime||0,e=this.audio.duration,s=t/e,o=this.state.playing;o=!!(s!=1&&o),this.setState({played:s,playing:o}),this.progressTimeout=setTimeout(this.progress,this.props.progressInterval/this.state.playbackRate)}},a.prototype.audioRef=function(t){this.audio=t},a.prototype.load=function(){this.setState({isReady:!0})},a.prototype.handlePlaybackRate=function(t){this.audio.playbackRate=t,this.setState({playbackRate:t,showHandlePlaybackRate:!1})},a.prototype.handleMute=function(){if(this.state.src){var t=this.state,e=t.muted,s=t.prevVolume,o=e?s:0;this.audio.muted=!e,this.setState({muted:!e,volume:o})}},a.prototype.handlePlaying=function(){if(this.state.src){var t=this.state.playing;t?this.audio.pause():this.audio.play(),this.setState({playing:!t})}},a.prototype.getCurrentTime=function(){if(!this.audio||!this.state.src||!this.state.isReady)return"0:00";var t=this.audio.duration,e=this.state.played;return this.formatTime(t*(e||0))},a.prototype.getDuration=function(){if(!this.audio||!this.state.src)return"0:00";if(!this.state.isReady)return this.onDurationCheck(),"0:00";var t=this.audio,e=t.duration,s=t.seekable;return e===1/0&&s.length>0?s.end(s.length-1):this.formatTime(e)},a.prototype.onDurationCheck=function(){clearTimeout(this.durationTimeout);var t=this.audio&&this.audio.duration;t||(this.durationTimeout=setTimeout(this.onDurationCheck,500))},a.prototype.onSeekChange=function(t){if(this.state.src){var e=t.target.value;this.setState({played:e})}},a.prototype.onSeekMouseDown=function(){this.setState({seeking:!0})},a.prototype.onSeekMouseUp=function(t){if(this.state.src&&this.state.seeking){var e=t.target.value,s=this.audio.duration;this.audio.currentTime=s*e;var o=this.props.loop,l=this.state.playing;l=e<1||o?l:!1,this.setState({playing:l,seeking:!1})}},a.prototype.setVolume=function(t){if(this.state.src){var e=t.target.value;this.audio.volume=e,this.setState({volume:e,prevVolume:e})}},a.prototype.formatTime=function(t){var e=new Date(t*1e3),s=e.getUTCHours(),o=isNaN(e.getUTCMinutes())?0:e.getUTCMinutes(),l=isNaN(e.getUTCSeconds())?"00":this.pad(e.getUTCSeconds());return s?"".concat(s,":").concat(this.pad(o),":").concat(l):"".concat(o,":").concat(l)},a.prototype.pad=function(t){return("0"+t).slice(-2)},a.prototype.toggleHandlePlaybackRate=function(){this.state.src&&this.setState({showHandlePlaybackRate:!this.state.showHandlePlaybackRate})},a.prototype.toggleHandleVolume=function(t){this.state.src&&this.setState({showHandleVolume:t})},a.prototype.renderRates=function(){var t=this,e=this.props,s=e.rates,o=e.classnames,l=this.state,m=l.showHandlePlaybackRate,h=l.playbackRate;return s&&s.length?m?i.createElement("div",{className:o("Audio-rateControl")},s.map(function(c,y){return i.createElement("div",{key:y,className:o("Audio-rateControlItem"),onClick:function(){return t.handlePlaybackRate(c)}},"x",c.toFixed(1))})):i.createElement("div",{className:o("Audio-rates"),onClick:this.toggleHandlePlaybackRate},"x",h.toFixed(1)):null},a.prototype.renderPlay=function(){var t=this.props.classnames,e=this.state.playing;return i.createElement("div",{className:t("Audio-play"),onClick:this.handlePlaying},e?i.createElement(d,{icon:"pause",className:"icon"}):i.createElement(d,{icon:"play",className:"icon"}))},a.prototype.renderTime=function(){var t=this.props.classnames;return i.createElement("div",{className:t("Audio-times")},this.getCurrentTime()," / ",this.getDuration())},a.prototype.renderProcess=function(){var t=this.props.classnames,e=this.state.played;return i.createElement("div",{className:t("Audio-process")},i.createElement("input",{type:"range",min:0,max:1,step:"any",value:e||0,onMouseDown:this.onSeekMouseDown,onChange:this.onSeekChange,onMouseUp:this.onSeekMouseUp}))},a.prototype.renderVolume=function(){var t=this,e=this.props.classnames,s=this.state,o=s.volume,l=s.showHandleVolume;return l?i.createElement("div",{className:e("Audio-volumeControl"),onMouseLeave:function(){return t.toggleHandleVolume(!1)}},i.createElement("div",{className:e("Audio-volumeControlIcon"),onClick:this.handleMute},o>0?i.createElement(d,{icon:"volume",className:"icon"}):i.createElement(d,{icon:"mute",className:"icon"})),i.createElement("input",{type:"range",min:0,max:1,step:"any",value:o,onChange:this.setVolume})):i.createElement("div",{className:e("Audio-volume"),onMouseEnter:function(){return t.toggleHandleVolume(!0)}},o>0?i.createElement(d,{icon:"volume",className:"icon"}):i.createElement(d,{icon:"mute",className:"icon"}))},a.prototype.render=function(){var t=this,e=this.props,s=e.className,o=e.style,l=e.inline,m=e.autoPlay,h=e.loop,c=e.controls,y=e.classnames,v=this.state,R=v.muted,b=v.src;return i.createElement("div",{className:y("Audio",s,l?"Audio--inline":""),style:o},i.createElement("audio",{className:y("Audio-original"),ref:this.audioRef,onCanPlay:this.load,autoPlay:m,controls:!0,muted:R,loop:h},i.createElement("source",{src:b})),i.createElement("div",{className:y("Audio-controls")},c&&c.map(function(g,C){g="render"+E(g);var P=g;return i.createElement(i.Fragment,{key:C},t[P]())})))},a.defaultProps={inline:!0,autoPlay:!1,playbackRate:1,loop:!1,rates:[],progressInterval:1e3,controls:["rates","play","time","process","volume"]},r([u,n("design:type",Function),n("design:paramtypes",[]),n("design:returntype",void 0)],a.prototype,"progress",null),r([u,n("design:type",Function),n("design:paramtypes",[HTMLMediaElement]),n("design:returntype",void 0)],a.prototype,"audioRef",null),r([u,n("design:type",Function),n("design:paramtypes",[]),n("design:returntype",void 0)],a.prototype,"load",null),r([u,n("design:type",Function),n("design:paramtypes",[Number]),n("design:returntype",void 0)],a.prototype,"handlePlaybackRate",null),r([u,n("design:type",Function),n("design:paramtypes",[]),n("design:returntype",void 0)],a.prototype,"handleMute",null),r([u,n("design:type",Function),n("design:paramtypes",[]),n("design:returntype",void 0)],a.prototype,"handlePlaying",null),r([u,n("design:type",Function),n("design:paramtypes",[]),n("design:returntype",void 0)],a.prototype,"getCurrentTime",null),r([u,n("design:type",Function),n("design:paramtypes",[]),n("design:returntype",void 0)],a.prototype,"getDuration",null),r([u,n("design:type",Function),n("design:paramtypes",[]),n("design:returntype",void 0)],a.prototype,"onDurationCheck",null),r([u,n("design:type",Function),n("design:paramtypes",[Object]),n("design:returntype",void 0)],a.prototype,"onSeekChange",null),r([u,n("design:type",Function),n("design:paramtypes",[]),n("design:returntype",void 0)],a.prototype,"onSeekMouseDown",null),r([u,n("design:type",Function),n("design:paramtypes",[Object]),n("design:returntype",void 0)],a.prototype,"onSeekMouseUp",null),r([u,n("design:type",Function),n("design:paramtypes",[Object]),n("design:returntype",void 0)],a.prototype,"setVolume",null),r([u,n("design:type",Function),n("design:paramtypes",[Number]),n("design:returntype",void 0)],a.prototype,"formatTime",null),r([u,n("design:type",Function),n("design:paramtypes",[Number]),n("design:returntype",void 0)],a.prototype,"pad",null),r([u,n("design:type",Function),n("design:paramtypes",[]),n("design:returntype",void 0)],a.prototype,"toggleHandlePlaybackRate",null),r([u,n("design:type",Function),n("design:paramtypes",[Boolean]),n("design:returntype",void 0)],a.prototype,"toggleHandleVolume",null),a}(i.Component),V=function(p){k(a,p);function a(){return p!==null&&p.apply(this,arguments)||this}return a=r([S({type:"audio"})],a),a}(F);export{F as Audio,V as AudioRenderer};