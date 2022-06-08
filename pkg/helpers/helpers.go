package helpers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandomStr(strlen int64) string {
	full := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLOMNOPQRSTUVWXYZ")
	fullLen := len(full) - 1
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]rune, 0)
	for i := 0; i <= int(strlen); i++ {
		result = append(result, full[myRand.Intn(fullLen)])
	}
	return string(result)
}

func GetLocaltion() *time.Location {
	l, _ := time.LoadLocation("Asia/Shanghai")
	return l
}

func GetTimestamp() int64 {
	return time.Now().In(GetLocaltion()).Unix()
}

func FillHttpScheme(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}
	if strings.HasPrefix(url, "https://") {
		return url
	}
	// 默认给https
	return "https://" + url
}

// convert a number to a base62 string
func ConvertNumToBase62(num uint) string {
	// base62 charactors
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	var digitals []uint
	for num != 0 {
		mod := num % 62
		digitals = append(digitals, mod)
		num = num / 62
	}

	reverse := func(arr []uint) []uint {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		return arr
	}
	digitals = reverse(digitals)
	var resultarr []string
	for _, elem := range digitals {
		resultarr = append(resultarr, chars[elem])
	}

	return strings.Join(resultarr, "")
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}
