/**
 * @param {number} n
 * @return {number}
 */
var subtractProductAndSum = function(n) {
    let x =String(n).split("")
    let p = 1
    for(let i of x){
        p*=Number(i)
    }
    let sum = 0
    for(let i of x){
        sum+=Number(i)
    }
    return p-sum
};