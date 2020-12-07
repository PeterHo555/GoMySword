package main  //声明main包

import (
	"fmt"
)

func main() {
	//fmt.Println(8)
	var a []int = []int{1,2,3,4,4}
	var ans int = findRepeatNumber(a)
	fmt.Println(ans)

}
