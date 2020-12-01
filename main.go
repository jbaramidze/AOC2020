package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readWords() []int {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	words := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, e := strconv.Atoi(scanner.Text())
		if e != nil {
			log.Fatal(e)
		}
		words = append(words, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

func d1_2() {
	m := make(map[int]bool)
	words := readWords()
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if _, ok := m[2020-words[i]-words[j]]; ok {
				log.Print(words[i] * words[j] * (2020 - words[i] - words[j]))
			}
		}
		m[words[i]] = true
	}
}

func d1_1() {
	m := make(map[int]bool)
	words := readWords()
	for i := 0; i < len(words); i++ {
		if _, ok := m[2020-words[i]]; ok {
			log.Print(words[i] * (2020 - words[i]))
		}
		m[words[i]] = true
	}
}

func main() {
	d1_1()
	d1_2()
}
