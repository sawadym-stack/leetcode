/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
    // let s= [...nums1,...nums2]
    // let newArr = s.filter(num => num !== 0);
    // let arr=newArr.sort((a,b)=>a-b)
    // console.log(arr)
    nums1.splice(m);
    nums1.push(...nums2)
    nums1.sort((a,b)=>a-b)
};