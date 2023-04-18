package utils

// Reversing a slice
func ReverseSlice(list []uint) {
    if len(list) > 1 {
        for i, j := 0, len(list) - 1; i < j; i, j = i + 1, j - 1 {
            list[i], list[j] = list[j], list[i]
        }
    }
}

func ReverseStringSlice(list []string) {
    if len(list) > 1 {
        for i, j := 0, len(list) - 1; i < j; i, j = i + 1, j - 1 {
            list[i], list[j] = list[j], list[i]
        }
    }
}
