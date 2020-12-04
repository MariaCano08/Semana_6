package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var s []string
	var c string
	var t int
	t = 1
	fileAscendente, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Printf("Fallo al abrir")
	}

	fileDescendente, err2 := os.Create("descendente.txt")
	if err2 != nil {
		fmt.Printf("Fallo al abrir el archivo descendente")
	}
	defer fileAscendente.Close()
	defer fileDescendente.Close()

	fmt.Printf("Catpure cada string que desea ordenar")
	fmt.Printf("Cuando termine de capturar escriba S")
	for i := 0; i < t; {
		fmt.Scan(&c)
		if "s" == strings.ToLower(c) {
			t = i + 2
			break
		}
		s = append(s, c)
		i = i + 1
		t = t + 1
	}
	//fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
	//fmt.Printf(s)

}
