package main

import "fmt"

func main() {

	var apontaus **int8
	var apontaus2 *int8
	var us int8 = 16
	apontaus2 = &us
	apontaus = &apontaus2
	fmt.Println(us, *apontaus2, **apontaus)

}
