package  homework

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
// 字符串数量
var wordCount = 50000000
// 字符串最大长度
var wordLength = 20
//每行字符串数量(每行大小)
var countInLine = 100

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func createWolrds() {
	// 1000万个字符串，长度1-20, 一行1000个 115.MB
	// 1亿个字符串，长度1-20, 一行1000个 1.15GB 58s

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