package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Pas entier")
		os.Exit(2)
	}
	fmt.Printf("%d\n", age)

	if age < 17 {
		fmt.Println("Trop jeune")
	} else {
		fmt.Println("Entrez :)")
	}
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(5) // 0 ou 1 ou 2 ou 3 ou 4
	if randInt == 0 {
		fmt.Println("Pas de chance")
	}

	switch {
	case age >= 100:
		fmt.Println("tres vieux")
	case age == 5 || age == 7:
		fmt.Println("Cool")
	default:
		fmt.Println("ok")
	}
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println(",,,,")
	case time.Thursday:
		fmt.Println("Hello there it's Thursday")
	default:
		fmt.Println("An other day")
	}
}
