package mapping

import (
	// "fmt"
	"goshima/utils"
	"strings"
    "math"
)

const BASE_62 uint = 62

type Charset struct {
	Start      rune
	Iterations int32
}

// Creates a bijective mapping between an integer and the corresponding string value
func EncodeMap(bijMap map[int32]string, charset Charset) {
	index := int32(len(bijMap))

	end := charset.Start + charset.Iterations

	for i := charset.Start; i <= end; i++ {
		bijMap[index] = string(i)
		index++
	}
}

// Convert primary key to shortened path
func GetShortPath(id uint, bijMap map[int32]string) string {
	digits := make([]uint, 0)

	var convertedString string

	for id > 0 {
		rem := id % BASE_62
		digits = append(digits, rem)
		id = id / BASE_62
	}

	utils.ReverseSlice(digits)

	for i := 0; i < len(digits); i++ {
		convertedString += bijMap[int32(digits[i])]
	}

	return convertedString
}

// Get primary key by converting shortPath
func GetUrlId(shortPath string, bijMap map[int32]string) uint {
    characters := strings.Split(shortPath, "")

    utils.ReverseStringSlice(characters)

    var primaryKey uint = 0

    for i := 0; i < len(characters); i++ {
        key, ok := GetKey(bijMap, characters[i])
        if !ok {
            panic("Error: Value does not exist in the Map")
        }

        decimalValue := math.Pow(float64(BASE_62), float64(i))
        primaryKey += uint(key * int32(decimalValue))
    }

    return primaryKey
}

// Helper function to get key from value
func GetKey(m map[int32]string, value string) (key int32, ok bool) {
    for k,v := range m {
        if v == value {
            key = k
            ok = true
            return
        }
    }

    return
}
