/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    // s=s.trim()
    // let words = s.split("")
    // let result = words[words.length-1]
    // return result.length
     s = s.trim();
     return s.length - s.lastIndexOf(' ') - 1;
};