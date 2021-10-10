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
}

func problem4() {
	// first marker
	const Line1 = 1.00
	// second marker
	line2 := 12.67
	distanceMeters := line2 - Line1
	if distanceMeters < 10.00 {
		fmt.Println("Unable to proceed, distance out of scope")
	}
	var car1 = 1.4579
	//d = st
	fmt.Println("Distance in Meters", distanceMeters)
	testSpeed := distanceMeters / car1
	fmt.Println("Test Speed in m/s", testSpeed)
	speedKMH := testSpeed * 18 / 5
	fmt.Println("in KM/H", speedKMH)
	//s = d/t
}

// is it worth it to speed?
func problem5() {
	totalDistance := 4500.0 // meters
	testSpeed := 23.0       // m/s
	//t = d/s
	time1 := totalDistance / testSpeed //in seconds
	time1Minutes := time1 / 60
	fmt.Println("Time in Seconds:", time1)
	fmt.Println("Time in Minutes:", time1Minutes)

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

func main() {
	problem1()
	problem2()
	problem3()
	problem4()
	problem5()
}
