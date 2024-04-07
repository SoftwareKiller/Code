package example

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/os/gtime"
)

func StrBuildCost() {
	strList := []string{"Hello", " ", "everyone", " !!"}

	times := 10
	begin := gtime.Millisecond()
	for i := 0; i < times; i++ {
		s := ""
		for _, str := range strList {
			s = fmt.Sprintf("%s%s", s, str)
		}
	}
	fmt.Printf("Case 1: Cost time:%d ms", gtime.Millisecond()-begin)

	for i := 0; i < times; i++ {
		_ = strings.Join(strList, "")
	}
	begin = gtime.Millisecond()
	fmt.Printf("Case 2:Cost time:%d ms", gtime.Millisecond()-begin)
}
