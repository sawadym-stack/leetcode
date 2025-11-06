/**
 * @param {string} s
 * @return {string}
 */
var clearDigits = function(s) {
    // if(s.includes(1||2||3||4||5||6||7||8||9||0)){
    //     return "";
    // }else{
    //     return s
    // }
  let arr = [];

  for (let i of s) {
    if (i >= '0' && i <= '9') {
      if (arr.length > 0) arr.pop();
    } else {
      arr.push(i);
    }
  }

  return arr.join('');
};