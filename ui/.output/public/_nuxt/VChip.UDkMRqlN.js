import{N as G,ac as Y,Q as z,bg as Z,V as L,W as R,X as T,Y as E,a0 as w,bh as H,a3 as J,a6 as r,a7 as ee,z as l,O as y,aT as B,P as ae,R as le,S as te,b_ as ne,T as se,by as ie,U as ce,ag as de,_ as oe,bA as ue,bB as re,ao as ve,bC as fe,bD as ke,b$ as pe,Z as me,c0 as be,bz as he,c as C,ai as _,aj as ye,bF as Ce,c1 as Ve,al as v,aI as f,bw as ge,ak as x,aG as D,a9 as Ie}from"./entry.Yh2omVp3.js";const F=Symbol.for("vuetify:v-chip-group"),Pe=G({column:Boolean,filter:Boolean,valueComparator:{type:Function,default:Y},...z(),...Z({selectedClass:"v-chip--selected"}),...L(),...R(),...T({variant:"tonal"})},"VChipGroup");E()({name:"VChipGroup",props:Pe(),emits:{"update:modelValue":e=>!0},setup(e,k){let{slots:c}=k;const{themeClasses:o}=w(e),{isSelected:t,select:p,next:m,prev:b,selected:h}=H(e,F);return J({VChip:{color:r(e,"color"),disabled:r(e,"disabled"),filter:r(e,"filter"),variant:r(e,"variant")}}),ee(()=>l(e.tag,{class:["v-chip-group",{"v-chip-group--column":e.column},o.value,e.class],style:e.style},{default:()=>{var u;return[(u=c.default)==null?void 0:u.call(c,{isSelected:t,select:p,next:m,prev:b,selected:h.value})]}})),{}}});const Se=G({activeClass:String,appendAvatar:String,appendIcon:y,closable:Boolean,closeIcon:{type:y,default:"$delete"},closeLabel:{type:String,default:"$vuetify.close"},draggable:Boolean,filter:Boolean,filterIcon:{type:String,default:"$complete"},label:Boolean,link:{type:Boolean,default:void 0},pill:Boolean,prependAvatar:String,prependIcon:y,ripple:{type:[Boolean,Object],default:!0},text:String,modelValue:{type:Boolean,default:!0},onClick:B(),onClickOnce:B(),...ae(),...z(),...le(),...te(),...ne(),...se(),...ie(),...ce(),...L({tag:"span"}),...R(),...T({variant:"tonal"})},"VChip"),Be=E()({name:"VChip",directives:{Ripple:de},props:Se(),emits:{"click:close":e=>!0,"update:modelValue":e=>!0,"group:selected":e=>!0,click:e=>!0},setup(e,k){let{attrs:c,emit:o,slots:t}=k;const{t:p}=oe(),{borderClasses:m}=ue(e),{colorClasses:b,colorStyles:h,variantClasses:u}=re(e),{densityClasses:O}=ve(e),{elevationClasses:$}=fe(e),{roundedClasses:M}=ke(e),{sizeClasses:j}=pe(e),{themeClasses:K}=w(e),V=me(e,"modelValue"),a=be(e,F,!1),s=he(e,c),N=C(()=>e.link!==!1&&s.isLink.value),i=C(()=>!e.disabled&&e.link!==!1&&(!!a||e.link||s.isClickable.value)),X=C(()=>({"aria-label":p(e.closeLabel),onClick(n){n.stopPropagation(),V.value=!1,o("click:close",n)}}));function g(n){var d;o("click",n),i.value&&((d=s.navigate)==null||d.call(s,n),a==null||a.toggle())}function q(n){(n.key==="Enter"||n.key===" ")&&(n.preventDefault(),g(n))}return()=>{const n=s.isLink.value?"a":e.tag,d=!!(e.appendIcon||e.appendAvatar),Q=!!(d||t.append),U=!!(t.close||e.closable),I=!!(t.filter||e.filter)&&a,P=!!(e.prependIcon||e.prependAvatar),W=!!(P||t.prepend),S=!a||a.isSelected.value;return V.value&&_(l(n,{class:["v-chip",{"v-chip--disabled":e.disabled,"v-chip--label":e.label,"v-chip--link":i.value,"v-chip--filter":I,"v-chip--pill":e.pill},K.value,m.value,S?b.value:void 0,O.value,$.value,M.value,j.value,u.value,a==null?void 0:a.selectedClass.value,e.class],style:[S?h.value:void 0,e.style],disabled:e.disabled||void 0,draggable:e.draggable,href:s.href.value,tabindex:i.value?0:void 0,onClick:g,onKeydown:i.value&&!N.value&&q},{default:()=>{var A;return[Ce(i.value,"v-chip"),I&&l(Ve,{key:"filter"},{default:()=>[_(l("div",{class:"v-chip__filter"},[t.filter?l(f,{key:"filter-defaults",disabled:!e.filterIcon,defaults:{VIcon:{icon:e.filterIcon}}},t.filter):l(v,{key:"filter-icon",icon:e.filterIcon},null)]),[[ge,a.isSelected.value]])]}),W&&l("div",{key:"prepend",class:"v-chip__prepend"},[t.prepend?l(f,{key:"prepend-defaults",disabled:!P,defaults:{VAvatar:{image:e.prependAvatar,start:!0},VIcon:{icon:e.prependIcon,start:!0}}},t.prepend):l(x,null,[e.prependIcon&&l(v,{key:"prepend-icon",icon:e.prependIcon,start:!0},null),e.prependAvatar&&l(D,{key:"prepend-avatar",image:e.prependAvatar,start:!0},null)])]),l("div",{class:"v-chip__content"},[((A=t.default)==null?void 0:A.call(t,{isSelected:a==null?void 0:a.isSelected.value,selectedClass:a==null?void 0:a.selectedClass.value,select:a==null?void 0:a.select,toggle:a==null?void 0:a.toggle,value:a==null?void 0:a.value.value,disabled:e.disabled}))??e.text]),Q&&l("div",{key:"append",class:"v-chip__append"},[t.append?l(f,{key:"append-defaults",disabled:!d,defaults:{VAvatar:{end:!0,image:e.appendAvatar},VIcon:{end:!0,icon:e.appendIcon}}},t.append):l(x,null,[e.appendIcon&&l(v,{key:"append-icon",end:!0,icon:e.appendIcon},null),e.appendAvatar&&l(D,{key:"append-avatar",end:!0,image:e.appendAvatar},null)])]),U&&l("button",Ie({key:"close",class:"v-chip__close",type:"button"},X.value),[t.close?l(f,{key:"close-defaults",defaults:{VIcon:{icon:e.closeIcon,size:"x-small"}}},t.close):l(v,{key:"close-icon",icon:e.closeIcon,size:"x-small"},null)])]}}),[[ye("ripple"),i.value&&e.ripple,null]])}}});export{Be as V};
