func reverse(x int) int {
    result := 0
    for x != 0 {
        digit := x % 10
        x /= 10

        if result > 2147483647/10 || (result == 2147483647/10 && digit > 7) {
            return 0
        }
        if result < -2147483648/10 || (result == -2147483648/10 && digit < -8) {
            return 0
        }

        result = result*10 + digit
    }
    return result
}