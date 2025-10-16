/**
 * @param {number[]} nums
 * @return {number}
 */
var differenceOfSum = function(nums) {
    let element=0;
    let digitSum =0;
    for(i=0;i<nums.length;i++){
        element+=nums[i];
    }
   for (let num of nums) {
  while (num > 0) {
    digitSum += num % 10; 
    num = Math.floor(num / 10); 
  }
}
return element-digitSum    
};