package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

var buf [1]byte

func readbyte(r io.Reader) (rune, error) {

	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {

	// For å generere rapporten om profil
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Kunne ikke åpne filen %q: %v", os.Args[1], err)
	}

	words := 0
	inword := false // tilstandsmaskin at vi er inn et ord (obs! unicode)
	b := bufio.NewReader(f)
	for {
		//r, err := readbyte(f) ->
		r, err := readbyte(b)
		if err == io.EOF {
			break
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}

	fmt.Printf("%q: %d words\n", os.Args[1], words)

}
