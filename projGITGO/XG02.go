package main

import "fmt"

func main() {

	var vuint8 string = "unsigned  8-bit integers (0 to 255)"
	var vuint16 string = "unsigned 16-bit integers (0 to 65535)"
	var vuint32 string = "unsigned 32-bit integers (0 to 4294967295)"
	var vuint64 string = "unsigned 64-bit integers (0 to 18446744073709551615)"
	var vint8 string = "signed  8-bit integers (-128 to 127)"
	var vint16 string = "signed 16-bit integers (-32768 to 32767)"
	var vint32 string = "signed 32-bit integers (-2147483648 to 2147483647)"
	var vint64 string = "signed 64-bit integers (-9223372036854775808 to 9223372036854775807)"

	fmt.Println(vuint8)
	fmt.Println(vuint16)
	fmt.Println(vuint32)
	fmt.Println(vuint64)
	fmt.Println(vint8)
	fmt.Println(vint16)
	fmt.Println(vint32)
	fmt.Println(vint64)

	cafe()
}

//Formatando strings para logs e mensagens

func cafe() {
	name := "Deg"
	coffee := 5

	fmt.Printf("Olá %s, você já bebeu %d cafés hoje", name, coffee)
}
