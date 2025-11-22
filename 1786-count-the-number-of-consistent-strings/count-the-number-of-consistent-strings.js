/**
 * @param {string} allowed
 * @param {string[]} words
 * @return {number}
 */
var countConsistentStrings = function(allowed, words) {
     const allowedSet = new Set(allowed);
    let count = 0;

    for (let word of words) {
        let isValid = true;

        for (let ch of word) {
            if (!allowedSet.has(ch)) {
                isValid = false;
                break;
            }
        }

        if (isValid) count++;
    }

    return count;
};