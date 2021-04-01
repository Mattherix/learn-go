package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, "Ah")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var age int

	for age < 18 {
		fmt.Println("Votre age: ")
		scanner.Scan()
		age, _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println("Bienvenue")

	for {
		fmt.Println("Boucle infni...")
		time.Sleep(5_000_000_000) // In nanosecond, Pourquoi ???
		// break
		// continue
	}
}
