package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{ //map from key strings to string slices
		"bond_james":      []string{"Shaken, not stirred", "martinis", "women"},
		"miss_moneypenny": []string{"james bond", "literature", "computer science"},
	}

	for k, v := range m { //Print out values, along with the index position of slices

		fmt.Printf("Key value: %v\n", k)
		for i, va := range v {
			fmt.Printf("Index %v\t, %v\n", i, va)
		}
	}
}
