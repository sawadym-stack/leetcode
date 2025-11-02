/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    if (n <= 2) return n;

  let first = 1; 
  let second = 2;
  let total=0;

  for (let i = 3; i <= n; i++) {
    total = first + second; 
    first = second;
    second = total;
  }

  return total;
};