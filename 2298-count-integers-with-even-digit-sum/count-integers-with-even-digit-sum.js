/**
 * @param {number} num
 * @return {number}
 */
var countEven = function (num) {
    let count = 0;

    for (let i = 2; i <= num; i++) {
        let sum = 0;
        let nm = i
        while (nm > 0) {
            let digit = nm % 10;
            sum += digit;
            nm = Math.floor(nm / 10);
        }
        if (sum % 2 == 0) {
            count++;
        }
    }
    return count
};