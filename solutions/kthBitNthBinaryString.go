package solutions

func findKthBit(n int, k int) byte {
    if n == 1 {
        return '0'
    }
    length := findLength(n)
    if k-1 > (length/2) {
        return inverse(findKthBit(n-1, length-k+1))
    }else if k-1 == length/2 {
        return '1'
    }else{
        return findKthBit(n-1, k)
    }
}

func findLength(n int) int {
    res := 1
    for i := 0; i < n-1; i++ {
        res = 2*res + 1
    }
    return res
}

func inverse(s byte) byte {
    if s == '0'{
        return '1'
    }
    return '0'
}