func myAtoi(s string) int {
     i := 0
    sign := 1
    result := 0

    const INT_MAX = 1<<31 - 1
    const INT_MIN = -1 << 31

    for i < len(s) && s[i] == ' ' {
        i++
    }

    if i < len(s) && (s[i] == '+' || s[i] == '-') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }

    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')

        if result > (INT_MAX-digit)/10 {
            if sign == 1 {
                return INT_MAX
            }
            return INT_MIN
        }

        result = result*10 + digit
        i++
    }

    return sign * result
}