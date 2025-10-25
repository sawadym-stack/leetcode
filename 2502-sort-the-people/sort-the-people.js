/**
 * @param {string[]} names
 * @param {number[]} heights
 * @return {string[]}
 */
var sortPeople = function(names, heights) {

let newf=names.map((value,index)=>[value,heights[index]])
newf.sort((a,b)=>b[1]-a[1])

console.log(newf)
let ok=newf.map((value)=>value[0])
return ok    
};