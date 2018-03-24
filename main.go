package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/suppayami/anytextify/services"
)

func read(scanner *bufio.Scanner) <-chan string {
	out := make(chan string)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			out <- text
		}
		close(out)
	}()
	return out
}

func transform(input <-chan string, service services.Textifier) <-chan string {
	out := make(chan string)
	go func() {
		for text := range input {
			out <- service.Transform(text)
		}
		close(out)
	}()
	return out
}

func handle(filename string, service services.Textifier) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := read(scanner)
	for text := range transform(input, service) {
		fmt.Println(text)
	}
}

func main() {
	var service services.Textifier

	filename := flag.String("f", "", "Input Text File")
	serviceType := flag.String("s", "", "Textify service")

	flag.Parse()

	if *filename == "" || *serviceType == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	service = services.GetService(*serviceType)
	handle(*filename, service)
}
