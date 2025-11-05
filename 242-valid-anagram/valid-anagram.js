/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    // let arr = [...s]
    // let arr1 =[...t]
    // for(let i = 0;i<arr.length;i++){
        
    // }
    // if(arr[i]==arr1[i]){
    //     return true
    // }else{
    //     false
    // }
    // console.log(arr)
     if (s.length !== t.length){
        return false;
     } 
    return s.split('').sort().join('') === t.split('').sort().join('');
};