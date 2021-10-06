package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
)

func problem1() {
	counts := make(map[string]int)
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Sorry, cannot find file to open")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		//counts[scanner.Text()]++
		counts[scanner.Text()]++
	}
	if 
	fmt.Println("Counts:", counts)
}

func problem2() {
	values := []byte("a friendly doggy! Woof, Woof, Woof, Grrr")
	result := bytes.Index(values, []byte("Woof"))
	fmt.Println("\n", result)
}

func problem3() {
	// Search for James Turnbull
	inputReturnLast := "Turnbull"
	//inputReturnFirst := "James"
	compileTrue := regexp.MustCompile(inputReturnLast)
	fmt.Println(compileTrue)
	//compileTrue.FindAllString("Tress, Teatro, Tannin, Telly, Turnbull, Turnbully, Tuddle", 1)
	fmt.Println(compileTrue.FindAllString("Tress, Teatro, Tannin, Telly, Turnbull, Turnbully, Tuddle", 1))
}

func problem4() {

}

func main() {
	problem1()
	problem2()
	problem3()
}
