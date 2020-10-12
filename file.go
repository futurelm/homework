package main

import (
	"bufio"
	"bytes"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"strconv"
	"strings"
)

// splitHugeFileToN split a huge file to N parts, return files name
func splitHugeFileToN(filePath string, N int) []fileWrite {
	var count = 0
	var wordChan = make(chan word)
	var subFiles = make([]fileWrite, count)
	var h = fnv.New64()

	go readFileToWord(filePath, wordChan)
	// create sub file
	for {
		fileName := fmt.Sprintf("subfile_%d.txt", count)
		f, err := os.Create(fileName)
		if err != nil {
			log.Printf("create file  %s failed", fileName)
			continue
		}
		// buffer = 1Gb
		subFiles = append(subFiles, fileWrite{f, bufio.NewWriterSize(f, 1024*1024*1024)})
		count++
		if count > N-1 {
			break
		}
	}
	for i := range wordChan {
		hash, _ := h.Write(i.value)
		//write string|index word to file
		_, err := subFiles[hash%N].wr.WriteString(fmt.Sprintf("%s|%d\n", string(i.value), i.index))
		if err != nil {
			log.Printf("str: %s write failed", string(i.value))
		}
	}

	for i := range subFiles {
		if err := subFiles[i].wr.Flush(); err != nil {
			log.Printf("file: %s last flush failed", subFiles[i].File.Name())
		}
	}
	return subFiles
}

//readFileToWord read file and send word to chan
func readFileToWord(filePath string, wordChan chan word) {
	var index int64
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		close(wordChan)
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for _, value := range bytes.Split(scanner.Bytes(), []byte(",")) {
			wordChan <- word{
				value: value,
				index: index,
			}
			index++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//readFileToWordMap read file and return a map with wordMeta
func readFileToWordMap(filePath string) map[string]wordMeta {
	var filter = make(map[string]wordMeta)
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := strings.Split(string(scanner.Bytes()), "|")
		ind, err := strconv.ParseInt(value[1], 10, 64)
		if err != nil {
			log.Print(err)
		}
		log.Printf("value[0]: %s , value[1]: %s , int value[1]:%d", value[0], value[1], ind)
		if v, ok := filter[value[0]]; ok {
			filter[value[0]] = wordMeta{value[1], v.count + 1}
		} else {
			filter[value[0]] = wordMeta{
				index: value[1],
				count: 1,
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return filter
}
