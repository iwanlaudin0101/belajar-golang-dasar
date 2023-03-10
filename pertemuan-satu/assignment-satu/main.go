package main

import "fmt"

func main() {
	i := 21
	fmt.Println(i)
	fmt.Printf("%T \n", i)

	fmt.Printf("%%\n")

	j := true
	fmt.Println(j)
	j = false
	fmt.Println(j)

	unicodeRusia := "\u042F"
	fmt.Println(unicodeRusia)

	base10, base8, base16f, base16F := 21, 25, 0xf, 0xF

	fmt.Printf("%d\n", base10)
	fmt.Printf("%o\n", base8)
	fmt.Printf("%x\n", base16f)
	fmt.Printf("%X\n", base16F)

	codeUnicodeRusia := 'Я'
	fmt.Printf("%U\n", codeUnicodeRusia)

	var k float64 = 123.456
	fmt.Println(k)
	fmt.Printf("%9f\n", k)
	fmt.Printf("%.6E", k)
}