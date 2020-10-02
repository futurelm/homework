package  homework

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
// words total count
var wordCount = 50000000
//  max length for one word
var wordLength = 20
// words total count in line
var countInLine = 100

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func createWolrds() {

	file, err := os.OpenFile("words.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)

	for i:=1; i<= wordCount; i++ {
		datawriter.WriteString(randSeq(1 + rand.Intn(wordLength))+",")
		if i% countInLine == 0 {
			datawriter.WriteString("\n")
			datawriter.Flush()
		}
	}
	file.Close()
}