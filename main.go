package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func readStrings() []string {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	words := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
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

func d2_1() {
	words := readStrings()
	valid := 0
	for _, w := range words {
		var min, max int
		var check rune
		var s string
		fmt.Sscanf(w, "%d-%d %c: %s", &min, &max, &check, &s)

		count := 0
		for _, c := range s {
			if c == check {
				count++
			}
		}

		if count >= min && count <= max {
			valid++
		}
	}

	log.Print(valid)
}

func d2_2() {
	words := readStrings()
	valid := 0
	for _, w := range words {
		var a, b int
		var check rune
		var s string
		fmt.Sscanf(w, "%d-%d %c: %s", &a, &b, &check, &s)

		runes := []rune(s)

		count := 0
		if runes[a-1] == check {
			count++
		}
		if runes[b-1] == check {
			count++
		}

		if count == 1 {
			valid++
		}
	}

	log.Print(valid)
}

func d3_1() {
	words := readStrings()
	x := 0
	count := 0
	for i := 1; i < len(words); i++ {
		x += 3
		if words[i][x%len(words[i])] == byte('#') {
			count++
		}
	}

	log.Println(count)
}

type d3Strategy struct {
	r int
	d int
}

func d3_2() {
	words := readStrings()
	strategies := []d3Strategy{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	prod := 1
	for _, s := range strategies {
		x := 0
		count := 0
		for i := s.d; i < len(words); i += s.d {
			x += s.r
			if words[i][x%len(words[i])] == byte('#') {
				count++
			}
		}

		prod = prod * count
	}

	log.Println(prod)
}

func d4_1() {
	words := readStrings()
	mustHaveFields := map[string]int{
		"byr": 1,
		"iyr": 1,
		"eyr": 1,
		"hgt": 1,
		"hcl": 1,
		"ecl": 1,
		"pid": 1,
	}

	count := 0

	for i := 0; i < len(words); i++ {
		l := ""
		for ; i < len(words) && len(words[i]) != 0; i++ {
			l = l + words[i] + " "
		}

		runes := []rune(l)

		ic := 0
		for len(runes) > 0 {
			var k string
			fmt.Sscanf(string(runes), "%s", &k)
			s := strings.Split(k, ":")[0]
			if mustHaveFields[s] == 1 {
				ic++
			}
			runes = runes[len(k)+1:]
		}
		if ic == 7 {
			count++
		}
	}

	log.Print(count)
}

func d4_2() {
	words := readStrings()
	mustHaveFields := map[string]interface{}{
		"byr": func(s string) bool {
			i, e := strconv.Atoi(s)
			return e == nil && i >= 1920 && i <= 2002
		},
		"iyr": func(s string) bool {
			i, e := strconv.Atoi(s)
			return e == nil && i >= 2010 && i <= 2020
		},
		"eyr": func(s string) bool {
			i, e := strconv.Atoi(s)
			return e == nil && i >= 2020 && i <= 2030
		},
		"hgt": func(s string) bool {
			i, e := strconv.Atoi(s[:len(s)-2])
			if e != nil {
				return false
			}
			if s[len(s)-2:] == "cm" {
				return i >= 150 && i <= 193
			}
			if s[len(s)-2:] == "in" {
				return i >= 59 && i <= 76
			}
			return false
		},
		"hcl": func(s string) bool {
			if s[0] != '#' || len(s) != 7 {
				return false
			}
			for i := 1; i <= 6; i++ {
				if s[i] >= '0' && s[i] <= '9' {
					continue
				}
				if s[i] >= 'a' && s[i] <= 'f' {
					continue
				}
				return false
			}

			return true
		},
		"ecl": func(s string) bool {
			avail := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			for _, t := range avail {
				if t == s {
					return true
				}
			}

			return false
		},
		"pid": func(s string) bool {
			if len(s) != 9 {
				return false
			}

			for _, t := range s {
				if t >= '0' && t <= '9' {
					continue
				}
				return false
			}

			return true
		},
	}

	count := 0

	for i := 0; i < len(words); i++ {
		l := ""
		for ; i < len(words) && len(words[i]) != 0; i++ {
			l = l + words[i] + " "
		}

		runes := []rune(l)

		ic := 0
		for len(runes) > 0 {
			var k string
			fmt.Sscanf(string(runes), "%s", &k)
			s := strings.Split(k, ":")
			if _, ok := mustHaveFields[s[0]]; ok == true {
				if mustHaveFields[s[0]].(func(string) bool)(s[1]) == true {
					ic++
				}
			}
			runes = runes[len(k)+1:]
		}
		if ic == 7 {
			count++
		}
	}

	log.Print(count)
}

func main() {
	d4_2()
}
