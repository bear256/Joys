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
	for scanner.Scan() {
		text := scanner.Text()
		count := strings.Count(text, "UCanUup")
		fmt.Println(count, text)
		total += count
	}
	fmt.Println(total)
}
