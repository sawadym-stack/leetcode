/**
 * @param {number[]} nums
 * @return {number}
 */
var findGCD = function(nums) {
    let greatest=Math.max(...nums);
    let smallest=Math.min(...nums);
    let arr=[];
    
          
        for(i=1;i<=smallest;i++){
        if(smallest%i==0&&greatest%i==0){
            arr.push(i)
        }
    
        
    }
    return Math.max(...arr)
    }


