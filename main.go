package main

import (
    "fmt"
)

func combinacoesDe(tamanhoComb int, elementos []int) [][]int {
    if tamanhoComb == 0 {
	return [][]int{{}}
    }
    if len(elementos) == 0 {
	return [][]int{}
    }
    primeiroElemento := elementos[0]
    restoElementos := elementos[1:]
    combinacoes1 := combinacoesDe(tamanhoComb-1, restoElementos)
    for i := 0; i < len(combinacoes1); i++ {
	combinacoes1[i] = append([]int{primeiroElemento}, combinacoes1[i]...)
    }
    combinacoes2 := combinacoesDe(tamanhoComb, restoElementos)
    return append(combinacoes1, combinacoes2...)
}

func main() {
    tamanhoComb := 5
    elementos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    resultado := combinacoesDe(tamanhoComb, elementos)

    for _, comb := range resultado {
	fmt.Println(comb)
    }
}
