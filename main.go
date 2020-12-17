package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
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

func d9_1() {
	words := readWords()
	preamble := 25
	for i := preamble; i < len(words); i++ {
		M := make(map[int]bool)
		found := false
		s := words[i]
		for j := i - preamble; j < i; j++ {
			if _, ok := M[s-words[j]]; ok {
				found = true
				break
			}
			M[words[j]] = true
		}

		if !found {
			log.Print(words[i])
		}
	}
}

func d9find(words []int) int {
	preamble := 25
	for i := preamble; i < len(words); i++ {
		M := make(map[int]bool)
		found := false
		s := words[i]
		for j := i - preamble; j < i; j++ {
			if _, ok := M[s-words[j]]; ok {
				found = true
				break
			}
			M[words[j]] = true
		}

		if !found {
			return words[i]
		}
	}

	return -1
}

func d9Min(r []int) int {
	m := r[0]
	for _, s := range r {
		if s < m {
			m = s
		}
	}
	return m
}

func d9Max(r []int) int {
	m := r[0]
	for _, s := range r {
		if s > m {
			m = s
		}
	}
	return m
}

func d9_2() {
	words := readWords()
	found := d9find(words)
	i := 0
	j := 0
	sum := words[0]
	for {
		if sum == found {
			print(d9Min(words[i:j+1]) + d9Max(words[i:j+1]))
			return
		} else if sum < found {
			j++
			sum = sum + words[j]
		} else if sum > found {
			sum = sum - words[i]
			i++
		}
	}
}

func d10_1() {
	words := readWords()
	sort.Ints(words)
	d := make([]int, 4)
	last := 0
	for _, w := range words {
		d[w-last]++
		last = w
	}

	log.Print(d[1], d[2], d[3]+1)
}

func d10_2() {
	words := readWords()
	sort.Ints(words)
	till := words[len(words)-1]
	d := make([]int, till+10)
	d[0] = 1
	for _, w := range words {
		for i := w - 3; i < w; i++ {
			if i >= 0 {
				d[w] += d[i]
			}
		}
	}

	log.Print(d[till])
}

func d11Adj(w [][]rune, x, y int) int {
	c := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if x+i < 0 || x+i >= len(w[0]) ||
				y+j < 0 || y+j >= len(w) ||
				(i == 0 && j == 0) {
				continue
			}

			if w[j+y][x+i] == '#' {
				c++
			}
		}
	}
	return c
}

func d11_1() {
	strings := readStrings()
	words := make([][]rune, 0)
	for _, s := range strings {
		words = append(words, []rune(s))
	}

	changed := true
	for changed {
		dst := make([][]rune, len(words))
		for i := range words {
			dst[i] = make([]rune, len(words[i]))
			copy(dst[i], words[i])
		}
		changed = false
		for i := 0; i < len(words); i++ {
			for j := 0; j < len(words[0]); j++ {
				adjs := d11Adj(words, j, i)
				if dst[i][j] == 'L' && adjs == 0 {
					dst[i][j] = rune('#')
					changed = true
				} else if dst[i][j] == '#' && adjs >= 4 {
					dst[i][j] = rune('L')
					changed = true
				}
			}
		}

		words = dst
	}

	c := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[0]); j++ {
			if words[i][j] == '#' {
				c++
			}
		}
	}

	log.Print(c)
}

func d11Adj2(w [][]rune, x, y int) int {
	c := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := 1; ; k++ {
				if x+k*i < 0 || x+k*i >= len(w[0]) ||
					y+k*j < 0 || y+k*j >= len(w) ||
					(i == 0 && j == 0) {
					break
				}

				if w[y+k*j][x+k*i] == '#' {
					c++
					break
				} else if w[y+k*j][x+k*i] == 'L' {
					break
				}
			}
		}
	}
	return c
}

func d11_2() {
	strings := readStrings()
	words := make([][]rune, 0)
	for _, s := range strings {
		words = append(words, []rune(s))
	}

	changed := true
	for changed {
		dst := make([][]rune, len(words))
		for i := range words {
			dst[i] = make([]rune, len(words[i]))
			copy(dst[i], words[i])
		}
		changed = false
		for i := 0; i < len(words); i++ {
			for j := 0; j < len(words[0]); j++ {
				adjs := d11Adj2(words, j, i)
				if dst[i][j] == 'L' && adjs == 0 {
					dst[i][j] = rune('#')
					changed = true
				} else if dst[i][j] == '#' && adjs >= 5 {
					dst[i][j] = rune('L')
					changed = true
				}
			}
		}

		words = dst
	}

	c := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[0]); j++ {
			if words[i][j] == '#' {
				c++
			}
		}
	}

	log.Print(c)
}

func d12Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func d12_1() {
	words := readStrings()
	n, e, deg := 0, 0, 0
	for _, w := range words {
		i, err := strconv.Atoi(w[1:])
		if err != nil {
			log.Fatal("Failed parsing")
		}
		switch w[0] {
		case 'N':
			n += i
		case 'S':
			n -= i
		case 'E':
			e += i
		case 'W':
			e -= i
		case 'L':
			deg -= i
		case 'R':
			deg += i
		case 'F':
			{
				deg = deg % 360
				if deg < 0 {
					deg += 360
				}
				switch deg {
				case 0:
					e += i
				case 90:
					n -= i
				case 180:
					e -= i
				case 270:
					n += i
				}
			}

		}
	}

	log.Print(d12Abs(n) + d12Abs(e))
}

func d12Rotate(x, y, degrees int) (int, int) {
	degrees = degrees % 360
	if degrees < 0 {
		degrees += 360
	}

	if degrees == 0 {
		return x, y
	}

	return d12Rotate(y, -x, degrees-90)
}

func d12_2() {
	words := readStrings()
	n, e, wpN, wpE := 0, 0, 1, 10
	for _, w := range words {
		i, err := strconv.Atoi(w[1:])
		if err != nil {
			log.Fatal("Failed parsing")
		}
		switch w[0] {
		case 'N':
			wpN += i
		case 'S':
			wpN -= i
		case 'E':
			wpE += i
		case 'W':
			wpE -= i
		case 'L':
			wpE, wpN = d12Rotate(wpE, wpN, -i)
		case 'R':
			wpE, wpN = d12Rotate(wpE, wpN, i)
		case 'F':
			{
				n += i * wpN
				e += i * wpE
			}

		}
	}

	log.Print(d12Abs(n) + d12Abs(e))
}

func d13_1() {
	words := readStrings()
	ts, e := strconv.Atoi(words[0])
	if e != nil {
		log.Fatal("Failed parsing")
	}

	idstrings := strings.Split(words[1], ",")
	ids := []int{}
	for _, id := range idstrings {
		if id == "x" {
			continue
		}
		idd, e := strconv.Atoi(id)
		if e != nil {
			log.Fatal("Failed parsing", id)
		}
		ids = append(ids, idd)
	}

	delay := math.MaxInt32
	ans := -1
	for _, id := range ids {
		d := id - (ts % id)
		if d < delay {
			delay = d
			ans = id * delay
		}
	}

	log.Print(ans)
}

func d13_2() {
	words := readStrings()
	idstrings := strings.Split(words[1], ",")
	ids := []int{}
	for _, id := range idstrings {
		if id == "x" {
			ids = append(ids, -1)
			continue
		}
		idd, e := strconv.Atoi(id)
		if e != nil {
			log.Fatal("Failed parsing", id)
		}
		ids = append(ids, idd)
	}

	ans := ids[0]
	step := ids[0]
	for i := 1; i < len(ids); i++ {
		if ids[i] == -1 {
			continue
		}
		for j := 1; ; j++ {
			if (ans+step*j+i)%ids[i] == 0 {
				ans, step = ans+step*j, step*ids[i]
				break
			}
		}
	}

	log.Print(ans)
}

func d14_1() {
	words := readStrings()
	M := make(map[int]int)
	mask0 := 0
	mask1 := 1
	for _, w := range words {
		s := strings.Split(w, "=")
		s[0] = strings.Trim(s[0], " ")
		s[1] = strings.Trim(s[1], " ")
		if s[0] == "mask" {
			mask0 = 0
			mask1 = 0
			for _, t := range s[1] {
				mask0 = mask0 * 2
				mask1 = mask1 * 2
				if t == 'X' {
					continue
				} else if t == '1' {
					mask1++
				} else if t == '0' {
					mask0++
				}
			}
			mask0 = ^mask0
		} else {
			addr, e := strconv.Atoi(s[0][4 : len(s[0])-1])
			if e != nil {
				log.Fatal("Parsing failed")
			}
			val, e := strconv.Atoi(s[1])
			if e != nil {
				log.Fatal("Parsing failed")
			}

			val = val | mask1
			val = val & mask0
			M[addr] = val
		}
	}

	sum := 0
	for _, v := range M {
		sum += v
	}

	log.Print(sum)
}

func d14is1(addr int, index int) bool {
	return ((1 << index) & addr) > 0
}

func d14f(M map[int]int, N int, pos int, mask string, val int, addr int) {
	if pos == len(mask) {
		M[N] = val
		return
	}

	N *= 10
	if mask[pos] == 'X' {
		d14f(M, N, pos+1, mask, val, addr)
		d14f(M, N+1, pos+1, mask, val, addr)
	} else if mask[pos] == '0' {
		if d14is1(addr, len(mask)-pos-1) {
			d14f(M, N+1, pos+1, mask, val, addr)
		} else {
			d14f(M, N, pos+1, mask, val, addr)
		}
	} else {
		d14f(M, N+1, pos+1, mask, val, addr)
	}
}

func d14_2() {
	words := readStrings()
	M := make(map[int]int)
	mask := ""
	for _, w := range words {
		s := strings.Split(w, "=")
		s[0] = strings.Trim(s[0], " ")
		s[1] = strings.Trim(s[1], " ")
		if s[0] == "mask" {
			mask = s[1]
		} else {
			addr, e := strconv.Atoi(s[0][4 : len(s[0])-1])
			if e != nil {
				log.Fatal("Parsing failed")
			}
			val, e := strconv.Atoi(s[1])
			if e != nil {
				log.Fatal("Parsing failed")
			}

			d14f(M, 0, 0, mask, val, addr)
		}
	}

	sum := 0
	for _, v := range M {
		sum += v
	}

	log.Print(sum)
}

func d15_1() {
	line := strings.Split(readStrings()[0], ",")
	words := []int{}
	for _, l := range line {
		n, e := strconv.Atoi(l)
		if e != nil {
			log.Fatal("Parsing failure")
		}
		words = append(words, n)
	}

	M := make(map[int][]int)
	M[0] = []int{}
	for i, w := range words {
		if _, ok := M[w]; !ok {
			M[w] = []int{}
		}

		M[w] = append(M[w], i)
	}

	last := words[len(words)-1]
	for i := len(M); i < 2020; i++ {
		var n int
		if len(M[last]) == 1 {
			n = 0
		} else {
			indexes := M[last][len(M[last])-2:]
			n = indexes[1] - indexes[0]
		}
		M[n] = append(M[n], i)
		last = n
	}

	log.Print(last)
}

func d15_2() {
	line := strings.Split(readStrings()[0], ",")
	words := []int{}
	for _, l := range line {
		n, e := strconv.Atoi(l)
		if e != nil {
			log.Fatal("Parsing failure")
		}
		words = append(words, n)
	}

	M := make(map[int][]int)
	M[0] = []int{}
	for i, w := range words {
		if _, ok := M[w]; !ok {
			M[w] = []int{}
		}

		M[w] = append(M[w], i)
	}

	last := words[len(words)-1]
	for i := len(M); i < 30000000; i++ {
		var n int
		if len(M[last]) == 1 {
			n = 0
		} else {
			indexes := M[last][len(M[last])-2:]
			n = indexes[1] - indexes[0]
		}
		M[n] = append(M[n], i)
		last = n
	}

	log.Print(last)
}

func d16_1() {
	words := readStrings()
	state := 0
	result := 0
	M := [][]int{}
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
			state++
			i++
		} else if state == 0 {
			var rgx = regexp.MustCompile(`(^[^:]*): ([\d]*)-([\d]*) or ([\d]*)-([\d]*)`)
			match := rgx.FindStringSubmatch(words[i])
			log.Print(match[1])
			a, _ := strconv.Atoi(match[2])
			b, _ := strconv.Atoi(match[3])
			c, _ := strconv.Atoi(match[4])
			d, _ := strconv.Atoi(match[5])
			M = append(M, []int{a, b})
			M = append(M, []int{c, d})
		} else if state == 2 {
			fields := strings.Split(words[i], ",")
			for _, f := range fields {
				i, e := strconv.Atoi(f)
				if e != nil {
					log.Fatal("Failed parsing")
				}
				found := false
				for _, m := range M {
					if m[0] <= i && i <= m[1] {
						found = true
						break
					}
				}
				if found == false {
					result += i
				}
			}
		}
	}

	log.Print(result)
}

type d16allowedRanges struct {
	name   string
	ranges [][]int
}

func d16f(perm []int, cache map[int][]int, permSet map[int]bool) []int {
	if len(perm) == len(cache) {
		return perm
	}

	idx := len(perm)

	for _, i := range cache[idx] {
		if permSet[i] {
			continue
		}

		permSet[i] = true
		resp := d16f(append(perm, i), cache, permSet)
		permSet[i] = false
		if len(resp) > 0 {
			return resp
		}
	}

	return []int{}
}

func d16_2() {
	words := readStrings()
	state := 0
	M := []d16allowedRanges{}
	validTickets := [][]int{}
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
			state++
			i++
		} else if state == 0 {
			var rgx = regexp.MustCompile(`(^[^:]*): ([\d]*)-([\d]*) or ([\d]*)-([\d]*)`)
			match := rgx.FindStringSubmatch(words[i])
			a, _ := strconv.Atoi(match[2])
			b, _ := strconv.Atoi(match[3])
			c, _ := strconv.Atoi(match[4])
			d, _ := strconv.Atoi(match[5])
			M = append(M, d16allowedRanges{name: match[1], ranges: [][]int{{a, b}, {c, d}}})
		} else if state == 1 {
			fields := strings.Split(words[i], ",")
			fieldInts := []int{}
			for _, f := range fields {
				i, e := strconv.Atoi(f)
				if e != nil {
					log.Fatal("Failed parsing")
				}
				fieldInts = append(fieldInts, i)
			}
			validTickets = append(validTickets, fieldInts)
		} else if state == 2 {
			fields := strings.Split(words[i], ",")
			fieldInts := []int{}
			valid := true
			for _, f := range fields {
				i, e := strconv.Atoi(f)
				if e != nil {
					log.Fatal("Failed parsing")
				}
				found := false
				for _, m := range M {
					if (m.ranges[0][0] <= i && i <= m.ranges[0][1]) ||
						(m.ranges[1][0] <= i && i <= m.ranges[1][1]) {
						found = true
						break
					}
				}
				if found == false {
					valid = false
					break
				}
				fieldInts = append(fieldInts, i)
			}
			if valid == true {
				validTickets = append(validTickets, fieldInts)
			}
		}
	}

	// allowedPositions[i] = allowed fields on position i
	allowedPositions := map[int][]int{}

	for i := 0; i < len(M); i++ {
		allowedPositions[i] = []int{}
		for j := 0; j < len(M); j++ {
			valid := true
			for _, t := range validTickets {
				if (t[i] < M[j].ranges[0][0] || t[i] > M[j].ranges[0][1]) &&
					(t[i] < M[j].ranges[1][0] || t[i] > M[j].ranges[1][1]) {
					valid = false
					break
				}
			}

			if valid {
				allowedPositions[i] = append(allowedPositions[i], j)
			}
		}
	}

	result := d16f([]int{}, allowedPositions, map[int]bool{})
	p := 1
	for i, r := range result {
		if strings.HasPrefix(M[r].name, "departure") {
			p = p * validTickets[0][i]
		}
	}

	log.Print(p)
}

func d17getNeighbors(M [][][]bool, z, x, y int) int {
	n := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}

				if i+z < 0 || j+x < 0 || k+y < 0 {
					continue
				}
				if i+z >= len(M) || j+x >= len(M) || k+y >= len(M) {
					continue
				}

				if M[i+z][j+x][k+y] {
					n++
				}

			}
		}
	}

	return n
}

func d17_1() {
	words := readStrings()
	offset := 30
	// z, y, x
	M1 := make([][][]bool, 2*offset+1)
	M2 := make([][][]bool, 2*offset+1)
	for i := 0; i < 2*offset+1; i++ {
		M1[i] = make([][]bool, 2*offset+1)
		M2[i] = make([][]bool, 2*offset+1)
		for j := 0; j < 2*offset+1; j++ {
			M1[i][j] = make([]bool, 2*offset+1)
			M2[i][j] = make([]bool, 2*offset+1)
		}
	}

	curr := M1
	prev := M2

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[0]); j++ {
			curr[0+offset][i+offset][j+offset] = words[i][j] == '#'
		}
	}

	totalActive := -1
	for c := 0; c < 6; c++ {
		curr, prev = prev, curr
		totalActive = 0
		for i := 0; i < 2*offset+1; i++ {
			for j := 0; j < 2*offset+1; j++ {
				for k := 0; k < 2*offset+1; k++ {
					active := prev[i][j][k]
					neighbors := d17getNeighbors(prev, i, j, k)
					if active {
						if neighbors == 2 || neighbors == 3 {
							curr[i][j][k] = true
							totalActive++
						} else {
							curr[i][j][k] = false
						}
					} else {
						if neighbors == 3 {
							curr[i][j][k] = true
							totalActive++
						} else {
							curr[i][j][k] = false
						}
					}
				}
			}
		}
	}

	log.Print(totalActive)
}

func d172getNeighbors(M [][][][]bool, w, z, x, y int) int {
	n := 0
	for t := -1; t <= 1; t++ {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					if t == 0 && i == 0 && j == 0 && k == 0 {
						continue
					}

					if t+w < 0 || i+z < 0 || j+x < 0 || k+y < 0 {
						continue
					}
					if t+w >= len(M) || i+z >= len(M) || j+x >= len(M) || k+y >= len(M) {
						continue
					}

					if M[t+w][i+z][j+x][k+y] {
						n++
					}

				}
			}
		}
	}

	return n
}

func d17_2() {
	words := readStrings()
	offset := 20
	// w, z, y, x
	M1 := make([][][][]bool, 2*offset+1)
	M2 := make([][][][]bool, 2*offset+1)
	for i := 0; i < 2*offset+1; i++ {
		M1[i] = make([][][]bool, 2*offset+1)
		M2[i] = make([][][]bool, 2*offset+1)
		for j := 0; j < 2*offset+1; j++ {
			M1[i][j] = make([][]bool, 2*offset+1)
			M2[i][j] = make([][]bool, 2*offset+1)
			for t := 0; t < 2*offset+1; t++ {
				M1[i][j][t] = make([]bool, 2*offset+1)
				M2[i][j][t] = make([]bool, 2*offset+1)
			}
		}
	}

	curr := M1
	prev := M2

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[0]); j++ {
			curr[0+offset][0+offset][i+offset][j+offset] = words[i][j] == '#'
		}
	}

	totalActive := -1
	for c := 0; c < 6; c++ {
		curr, prev = prev, curr
		totalActive = 0
		for t := 0; t < 2*offset+1; t++ {
			for i := 0; i < 2*offset+1; i++ {
				for j := 0; j < 2*offset+1; j++ {
					for k := 0; k < 2*offset+1; k++ {
						active := prev[t][i][j][k]
						neighbors := d172getNeighbors(prev, t, i, j, k)
						if active {
							if neighbors == 2 || neighbors == 3 {
								curr[t][i][j][k] = true
								totalActive++
							} else {
								curr[t][i][j][k] = false
							}
						} else {
							if neighbors == 3 {
								curr[t][i][j][k] = true
								totalActive++
							} else {
								curr[t][i][j][k] = false
							}
						}
					}
				}
			}
		}
	}

	log.Print(totalActive)

}

func main() {
	d17_2()
}
