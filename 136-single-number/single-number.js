/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
     const a = new Set(nums);
  let suma = 0, sumAll = 0;
  
  for (let n of a) suma += n;
  for (let n of nums) sumAll += n;
  
  return 2 * suma - sumAll;
};