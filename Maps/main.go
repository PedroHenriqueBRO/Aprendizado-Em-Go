package main

import "fmt"

func main() {

	mapeamento := map[string]int{
		"Pedro":    2004,
		"Joao":     2007,
		"Cabrunco": 2009,
	}
	fmt.Println(mapeamento["Pedro"])
	mapeamentomapeamento := map[string]map[int]string{
		"2000-2010": {
			2004: "Pedro,ana,tim,vinicius",
			2005: "Naosei,naoseitambem",
		},
	}
	fmt.Println(mapeamentomapeamento["2000-2010"][2004])
	//delete(mapeamentomapeamento, "2000-2010") deleta
	fmt.Println(mapeamentomapeamento)
	mapeamentomapeamento["2011-2020"] = map[int]string{
		2011: "Kiko,Chaves",
	}
	fmt.Println(mapeamentomapeamento["2011-2020"][2011])

}
