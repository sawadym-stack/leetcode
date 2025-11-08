/**
 * @param {number} num
 * @return {number}
 */
var addDigits = function(num) {
//     let arr = num.toString().split('').map(Number)
//     let sum = arr.reduce((a,b)=>a+b,0)
//     let dgt = sum.toString().split('').map(Number)
//     let sum2 = dgt.reduce((a,b)=>a+b,0)
// return sum2
//     console.log(sum2)

 while (num > 9) {
        num = num.toString().split('').map(Number).reduce((a, b) => a + b, 0);
    }
    return num;
};