package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/jerseybbq/choose"
)

type sortRune []byte

func (s sortRune) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRune) Len() int {
	return len(s)
}

func UniqueLetters(word string) sortRune {
	var letters sortRune
	for n := 0; n < len(word); n++ {
		if word[n] == '\'' {
			continue
		}

		gotIt := false
		for _, c := range letters {
			if word[n] == c {
				gotIt = true
				break
			}
		}
		if !gotIt {
			letters = append(letters, word[n])
		}
	}
	sort.Sort(sortRune(letters))
	return letters
}

func ParseDictionary(path string) map[string][]string {
	dict := make(map[string][]string)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		letters := string(UniqueLetters(word))
		if len(letters) < 5 || len(letters) > 7 {
			// game rules
			continue
		}
		dict[letters] = append(dict[letters], word)
	}
	return dict
}

func main() {
	var dictfile, outer, center string

	flag.StringVar(&dictfile, "dict", "", "dictionary file")
	flag.StringVar(&outer, "outer", "", "outer letters")
	flag.StringVar(&center, "center", "", "center letter")
	flag.Parse()

	if len(outer) != 6 || len(center) != 1 {
		flag.Usage()
		return
	}

	d := ParseDictionary(dictfile)

	// puzzle is only interested in words of length 5 to 7
	for k := 4; k < 7; k++ {
		combos := choose.ChooseString(outer, k)
		for _, c := range combos {
			c2 := sortRune(fmt.Sprintf("%s%s", c, center))
			sort.Sort(c2)
			results := d[string(c2)]
			for _, w := range results {
				fmt.Println(w)
			}
		}
	}
}
