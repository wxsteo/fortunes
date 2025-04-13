package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	file, err := os.Open("content.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	var lines []string

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		if len(line) < 1 {
			continue
		}

		lines = append(lines, string(line))
	}

	if len(lines) < 1 {
		panic("empty fortune file")
	}

	statement := rand.Intn(len(lines))

	fmt.Println(lines[statement])

}
