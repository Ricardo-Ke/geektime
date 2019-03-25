package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var builder1 strings.Builder
	// builder1.WriteString("lalalala")
	// builder1.WriteString(" haohaohao")
	// // fmt.Printf("%s\n", builder1.String())

	// fmt.Println(builder1.String())
	// // 省略若干代码。
	// fmt.Println("Grow the builder ...")
	// builder1.Grow(10)
	// fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())

	// str := "This is an example of a string"
	// fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	// fmt.Printf("%t\n", strings.HasSuffix(str, "ings"))
	// fmt.Printf("%t\n", strings.Contains(str, "a"))
	// fmt.Printf("%d\n", strings.Index(str, "an"))
	// fmt.Printf("%d\n", strings.LastIndex(str, "an"))
	// fmt.Printf("%d\n", strings.IndexRune(str, rune('s')))

	// var s string = "Hello, how is it going, Hugo?罗森"
	// var g string = "ggggggggggg"
	// fmt.Printf("%d\n", strings.Count(s, "H"))
	// fmt.Printf("%d\n", strings.Count(g, "gg"))
	// fmt.Printf("%s\n", strings.Repeat(g, 2))
	// fmt.Printf("%s\n", strings.ToLower(s))
	// fmt.Printf("%s\n", strings.ToUpper(s))

	// s := "The quick brown fox jumps over the lazy dog"
	// sl := strings.Fields(s)
	// fmt.Printf("%v\n", sl)
	// for i, val := range sl {
	// 	fmt.Printf("%d - %s\t\t", i, val)
	// }
	// fmt.Println()
	// s2 := "GO1|The ABC of Go|25"
	// sl2 := strings.Split(s2, "|")
	// for i, val := range sl2 {
	// 	fmt.Printf("%d - %s\t\t", i, val)
	// }
	// fmt.Println()
	// s3 := strings.Join(sl2, "-")
	// fmt.Print(s3)

	//==============================  strconv  ==============================
	fmt.Println(strconv.IntSize)

	var orig string = "666"
	var an int
	var news string

	an, _ = strconv.Atoi(orig)
	fmt.Printf("%d\n", an)
	an = an + 5
	news = strconv.Itoa(an)
	fmt.Printf("%s\n", news)
}
