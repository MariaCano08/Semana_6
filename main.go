package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	//"log"
)

func main() {
	var s []string
	var c string
	var t, i, j int
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
	fmt.Printf("Cuando termine de capturar escriba S: ")
	for i = 0; i < t; {
		fmt.Scan(&c)
		if "s" == strings.ToLower(c) {
			t = i + 2
			break
		}
		c = c + "\n"
		s = append(s, c)
		i = i + 1
		t = t + 1
	}

	sort.Strings(s)
	ss := ""
	ss2 := ""
	for j = 0; j < i; {
		ss += s[j]
		j = j + 1
	}
	j = j - 1
	for j >= 0 {
		ss2 += s[j]
		j = j - 1
	}

	b := []byte(ss)
	e := ioutil.WriteFile("ascendente.txt", b, 0644)
	if e != nil {
		fmt.Println(e)
		return
	}
	b = []byte(ss2)
	e = ioutil.WriteFile("descendente.txt", b, 0644)
	if e != nil {
		fmt.Println(e)
		return
	}

}
