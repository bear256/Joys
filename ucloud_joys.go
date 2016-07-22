package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

import (
	"fmt"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		count := strings.Count(text, "UCanUup")
		fmt.Println("Line", i, ":", count)
		total += count
	}
	fmt.Println("Total:", total)
}
