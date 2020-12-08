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
	mustHaveFields := map[string](func(string) bool){
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
				if mustHaveFields[s[0]](s[1]) == true {
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

func pow(x int64, y int) int64 {
	var r int64 = 1
	for i := 0; i < y; i++ {
		r *= x
	}

	return r
}

func d5_1() {
	words := readStrings()
	h := 0
	for _, w := range words {
		a := 0
		for i := 0; i < 7; i++ {
			if w[i] == 'B' {
				a += int(pow(2, 7-i-1))
			}
		}
		b := 0
		for i := 0; i < 3; i++ {
			if w[i+7] == 'R' {
				b += int(pow(2, 3-i-1))
			}
		}
		r := a*8 + b
		if r > h {
			h = r
		}
	}

	log.Print(h)
}

func d5_2() {
	words := readStrings()
	M := make(map[int]bool)
	for _, w := range words {
		a := 0
		for i := 0; i < 7; i++ {
			if w[i] == 'B' {
				a += int(pow(2, 7-i-1))
			}
		}
		b := 0
		for i := 0; i < 3; i++ {
			if w[i+7] == 'R' {
				b += int(pow(2, 3-i-1))
			}
		}
		r := a*8 + b
		M[r] = true
	}

	for i := 1; i < 947; i++ {
		_, o1 := M[i-1]
		_, o2 := M[i]
		_, o3 := M[i+1]

		if o1 && !o2 && o3 {
			log.Print(i)
		}
	}
}

func d6_1() {
	words := readStrings()
	result := 0
	for i := 0; i < len(words); i++ {
		m := make(map[rune]bool)
		for ; i < len(words) && len(words[i]) > 0; i++ {
			for _, k := range words[i] {
				m[k] = true
			}
		}
		for i := 'a'; i <= 'z'; i++ {
			if _, ok := m[i]; ok {
				result++
			}
		}
	}
	log.Print(result)
}

func d6_2() {
	words := readStrings()
	result := 0
	for i := 0; i < len(words); i++ {
		m := make(map[rune]int)
		c := 0
		for ; i < len(words) && len(words[i]) > 0; i++ {
			for _, k := range words[i] {
				if _, ok := m[k]; !ok {
					m[k] = 0
				}
				m[k]++
			}
			c++
		}
		for i := 'a'; i <= 'z'; i++ {
			if m[i] == c {
				result++
			}
		}
	}
	log.Print(result)
}

type d7struct struct {
	count int
	bag   string
}

func d7checkCanContain(M map[string][]d7struct, b string) bool {
	for _, val := range M[b] {
		if val.bag == "shiny gold" {
			return true
		}
		if d7checkCanContain(M, val.bag) {
			return true
		}
	}

	return false
}

func d7_1() {
	wordsArray := readStrings()

	srcToContent := make(map[string][]d7struct)
	dstToParent := make(map[string][]string)

	for _, words := range wordsArray {
		fields := strings.Fields(words)
		src := fields[0] + " " + fields[1]
		if _, ok := srcToContent[src]; !ok {
			srcToContent[src] = []d7struct{}
		}

		for i := 4; i < len(fields); {
			if fields[i] == "no" {
				break
			}
			n, e := strconv.Atoi(fields[i])
			if e != nil {
				log.Fatalf("FAILED PARSING %v", fields[i])
			}
			dst := fields[i+1] + " " + fields[i+2]
			srcToContent[src] = append(srcToContent[src], d7struct{n, dst})
			if _, ok := dstToParent[dst]; !ok {
				dstToParent[dst] = []string{}
			}
			dstToParent[dst] = append(dstToParent[dst], src)
			i += 4
		}
	}

	result := 0
	for k := range srcToContent {
		if d7checkCanContain(srcToContent, k) {
			result++
		}
	}

	log.Print(result)
}

func d7countBagsInside(M map[string][]d7struct, b string) int64 {
	var c int64 = 0
	for _, v := range M[b] {
		c += int64(v.count) + int64(v.count)*(d7countBagsInside(M, v.bag))
	}
	return c
}

func d7_2() {
	wordsArray := readStrings()

	srcToContent := make(map[string][]d7struct)
	dstToParent := make(map[string][]string)

	for _, words := range wordsArray {
		fields := strings.Fields(words)
		src := fields[0] + " " + fields[1]
		if _, ok := srcToContent[src]; !ok {
			srcToContent[src] = []d7struct{}
		}

		for i := 4; i < len(fields); {
			if fields[i] == "no" {
				break
			}
			n, e := strconv.Atoi(fields[i])
			if e != nil {
				log.Fatalf("FAILED PARSING %v", fields[i])
			}
			dst := fields[i+1] + " " + fields[i+2]
			srcToContent[src] = append(srcToContent[src], d7struct{n, dst})
			if _, ok := dstToParent[dst]; !ok {
				dstToParent[dst] = []string{}
			}
			dstToParent[dst] = append(dstToParent[dst], src)
			i += 4
		}
	}

	log.Print(d7countBagsInside(srcToContent, "shiny gold"))
}

func d8_1() {
	words := readStrings()
	M := make(map[int]bool)
	acc := 0
	i := 0
	for {
		instr := words[i][:3]
		n, e := strconv.Atoi(words[i][4:])
		if e != nil {
			log.Fatal("Failed parsing")
		}
		if _, ok := M[i]; ok {
			log.Print(acc)
			return
		}
		M[i] = true

		if instr == "nop" {
			i++
		} else if instr == "acc" {
			acc = acc + n
			i++
		} else {
			i = i + n
		}
	}
}

func d8terminates(words []string, M map[int]int, i int, acc int) int {
	for {
		if i == len(words) {
			return acc
		}
		instr := words[i][:3]
		n, e := strconv.Atoi(words[i][4:])
		if e != nil {
			log.Fatal("Failed parsing")
		}

		if v, ok := M[i]; ok || v == -2 {
			return -1
		}

		M[i] = -2 // Meaning that originally we cannot reach it, but if we somehow do, it's a dead end

		if instr == "nop" {
			i++
		} else if instr == "acc" {
			acc = acc + n
			i++
		} else {
			i = i + n
		}
	}
}

func d8_2() {
	words := readStrings()
	M := make(map[int]int)
	acc := 0
	i := 0
	for {
		instr := words[i][:3]
		n, e := strconv.Atoi(words[i][4:])
		if e != nil {
			log.Fatal("Failed parsing")
		}
		if _, ok := M[i]; ok {
			break
		}
		M[i] = acc

		if instr == "nop" {
			i++
		} else if instr == "acc" {
			acc = acc + n
			i++
		} else {
			i = i + n
		}
	}

	// In each M[i] we have acc at that day.
	for i := range words {
		if v, ok := M[i]; !ok || v == -2 {
			continue
		}
		instr := words[i][:3]
		n, e := strconv.Atoi(words[i][4:])
		if e != nil {
			log.Fatal("Failed parsing")
		}
		if instr == "acc" {
			continue
		} else {
			j := 0
			if instr == "nop" {
				// try jumping
				j = i + n
			} else {
				// do nop instead jump
				j = i + 1
			}

			if v := d8terminates(words, M, j, M[i]); v >= 0 {
				log.Print(v)
				return
			}
		}

	}
}

func main() {
	d8_2()
}
