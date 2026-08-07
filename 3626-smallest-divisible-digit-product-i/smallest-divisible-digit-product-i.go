func smallestNumber(n int, t int) int {
        for {
        if isValid(n, t) {
            return n
        }
        n++
    }
}

func isValid(num int, t int) bool {
    prod := 1
    for num > 0 {
        digit := num % 10
        
        // If digit is 0 → product becomes 0 → always divisible
        if digit == 0 {
            return true
        }
        
        prod *= digit
        num /= 10
    }
    
    return prod%t == 0
}