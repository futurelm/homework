package homework

import (
	"bufio"
	"bytes"
	"hash/fnv"
	"log"
	"math"
	"os"
	"runtime"
	"sync"
)

const  ChanArrSize  = 10
const  ChanBuffer = 5
const  SplitChar = ","

type word struct {
	value  string
	index  int64
}

func findNonRepeatedString(file string) {

	var chanArr = [ChanArrSize]chan word{}
	var firstNonRepeatedChan = make(chan word, ChanBuffer)
	for i := range chanArr {
		chanArr[i] = make(chan word,ChanBuffer)
	}

    go readFileLineByLineOutChan(file, chanArr)
	go readChanOutputNonRepeatedChan(chanArr,firstNonRepeatedChan)

	var minIndex int64 = math.MaxInt64
	var firstNonRepeatedstring string
	for nonRepeatedWord := range firstNonRepeatedChan {
		if minIndex > nonRepeatedWord.index {
			minIndex = nonRepeatedWord.index
			firstNonRepeatedstring = nonRepeatedWord.value
		}
	}
	log.Printf("first non-repeated string is %s, index is %d \n",firstNonRepeatedstring,minIndex)
}

func readFileLineByLineOutChan(filePath string, chanArr [ChanArrSize]chan word) {
	var index int64
	h:=fnv.New64()
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for _,value := range bytes.Split(scanner.Bytes(), []byte(",")) {
			hashValue, _ := h.Write(value)
			chanArr[hashValue % ChanArrSize] <- word{
				value: string(value),
				index: index,
			}
			index++
		}
	}


	for i := range chanArr {
		close(chanArr[i])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readChanOutputNonRepeatedChan(inChans [ChanArrSize]chan word, outChan chan word) {
	var wg sync.WaitGroup

	for i := range inChans {
		wg.Add(1)
		go func(i int) {
			filterInChan(inChans[i],outChan)
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(outChan)
}


func filterInChan(in,out chan word) {
	wordMap := make(map[string]int64)
	for word := range in {
		index, ok := wordMap[word.value]
		if ok && index != -1{
			wordMap[word.value] = -1
		} else {
			wordMap[word.value] = word.index
		}
	}
	for k,v := range wordMap {
		if v != -1 {
			out <- word{
				value: k,
				index: v,
			}
		}
	}
	runtime.GC()
}
