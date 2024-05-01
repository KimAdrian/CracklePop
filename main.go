package main

import (
	"fmt"
	"strconv"
)

type rules struct {
	key int
	val string
}

func main() {
	rulesStore := []rules{
		{3, "Crackle"},
		{5, "Pop"},
		{15, "CracklePop"},
	}

	for i := 1; i < 100; i++ {
		result := ""
		for _, rule := range rulesStore {
			if i%rule.key == 0 {
				result = rule.val
			}
		}
		if result == "" {
			result = strconv.Itoa(i)
		}
		fmt.Println(result)
	}
}
