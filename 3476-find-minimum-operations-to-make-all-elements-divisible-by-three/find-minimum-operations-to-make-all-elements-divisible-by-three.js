/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumOperations = function(nums) {
    let operation=0;
    for(let num of nums){
        const x = num % 3
        operation += Math.min(x,3-x)
    }
    return operation
};