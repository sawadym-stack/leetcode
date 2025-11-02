/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    let result = [];

  for (let str of s) {
    if (str === '(') result.push(')');
    else if (str === '{') result.push('}');
    else if (str === '[') result.push(']');
    else if (result.pop() !== str) return false;
  }

  return result.length === 0;
};