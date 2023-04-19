package main

import "fmt"

func main() {

	strLit := "strLit"
	var varStr = "varStr"
	const constStr = "constStr"
	const typedConstStr string = "typedConstStr"

	printTypeAndValue(strLit)
	printTypeAndValue(varStr)
	printTypeAndValue(constStr)
	printTypeAndValue(typedConstStr)

	type CustomString string

	var customString CustomString
	//customString = strLit //cant, need conversion like CustomString(strLit)
	//customString = varStr //cant
	//customString = typedConstStr //cant

	customString = constStr
	printTypeAndValue(customString)

}

func printTypeAndValue(value interface{}) {
	fmt.Printf("%v %T\n", value, value)
}
