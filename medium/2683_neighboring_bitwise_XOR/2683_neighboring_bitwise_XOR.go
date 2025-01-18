package neighboring_bitwise_XOR


func doesValidArrayExist(derived []int) bool {
    original := []int{0}
    for i := 0; i < len(derived); i++ {
        original = append(original, derived[i]^original[i])
    }
    checkForZero := original[0] == original[len(original)-1]

    original = []int{1}
    for i := 0; i < len(derived); i++ {
        original = append(original, derived[i]^original[i])
    }
    checkForOne := original[0] == original[len(original)-1]

    return checkForZero || checkForOne
}