// Tutoriel suivi -> https://devopssec.fr/article/cours-apprendre-langage-programmation-go#begin-article-section
package main

import (
	"fmt"
)

func main() {
	// Commentaire
	/* Commentaire
	multi-ligne
	*/
	fmt.Println("Hello world")

	/* Type de var:
	uint8, uint16, uint32, uint64 0->...
	int8, int16, int32, int64 -128, ... -> 127, ...
	float32, float64 ~7 ou 16 décimals
	byte == uint8
	rune == uint16
	uint and int si machine 32 ou 64 bits
	bool true or false
	string
	*/
	var test int = 16 // Surchargement de la variable
	var (
		vie int    = 1
		c   string = "Hello"
	)

	fmt.Printf("%d\n", test)
	fmt.Printf("%d %t %s\n", vie, false, c)

	// Typage dynamique
	kill := 5
	fmt.Printf("%d\n", kill)

	const maConstante = 500

	fmt.Printf("%d\n", maConstante)
	fmt.Printf("%T\n", false)

	a := 3
	b := 2

	fmt.Println("a + b  = ", a+b) // addition de la variable a et b
	fmt.Println("a - b  = ", a-b) // soustraction de la variable a et b
	fmt.Println("a * b  = ", a*b) // multiplication de la variable a et b
	fmt.Println("a / b  = ", a/b) // division de la variable a et b
	fmt.Println("a % b  = ", a%b) // modulo de la variable a et b

	// += -= *= /= %= et ++ --

	x := 3.5
	y := 3
	fmt.Printf("x + y = %f\n", x+float64(y))
}
