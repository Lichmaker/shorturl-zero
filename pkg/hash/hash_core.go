package hash

import (
	"github.com/lichmaker/short-url-micro/pkg/helpers"
	"github.com/spaolacci/murmur3"
)

func Make(long string) string {
	num := murmur3.Sum32([]byte(long))

	return helpers.ConvertNumToBase62(uint(num))
}
