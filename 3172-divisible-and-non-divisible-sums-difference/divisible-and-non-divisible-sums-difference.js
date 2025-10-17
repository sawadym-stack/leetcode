/**
 * @param {number} n
 * @param {number} m
 * @return {number}
 */
var differenceOfSums = function(n, m) {
    let divmsum = 0;
    let ndivmsum = 0;
    for(let i=1;i<=n;i++){
        if(i%m!==0){
          ndivmsum+=i;
        }else if(i%m==0){
            divmsum+=i;
        }
    }
    return ndivmsum-divmsum
};
