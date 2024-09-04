package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"unique"
)

// Before running this program, run ./generate.sh to
// produce the file Le_Comte_de_Monte-Cristo_x100.txt

func main() {
	const bookPath = "./Le_Comte_de_Monte-Cristo_x100.txt"
	data, err := os.ReadFile(bookPath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Read", len(data), "bytes from", bookPath)
	mem()
	book := string(data)
	Bwords := findBwords(book)
	mem()
	// Use Bwords
	fmt.Printf("The last B-word is %q\n", Bwords[len(Bwords)-1].Value())
	mem()
}

const wordchars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-àâéèëêëîôû"

var iswordchar = map[rune]bool{}

func init() {
	for _, c := range wordchars {
		iswordchar[c] = true
	}
}

func findBwords(book string) []unique.Handle[string] {
	n := 0
	var Bwords []unique.Handle[string]

	a := -1
	for i, c := range book {
		if iswordchar[c] {
			// current char is in a word e.g. 'a', 'à', 'm'
			if a == -1 {
				// start of a word
				a = i
			}
		} else {
			// current char is not in word e.g. ' ', ','
			if a != -1 {
				// just finished a word
				n++
				word := book[a:i]
				if word[0] == 'b' || word[0] == 'B' {
					cloned := strings.Clone(word)
					handle := unique.Make(cloned)
					Bwords = append(Bwords, handle)
				}
			}
			a = -1
		}
	}
	fmt.Println("Found", len(Bwords), "B-words out of", n, "words")
	return Bwords
}

var memstat runtime.MemStats

func mem() {
	runtime.GC()
	runtime.ReadMemStats(&memstat)
	const MiB = 1024 * 1024
	fmt.Println("The program is now using", memstat.Alloc/MiB, "MiB")
}
