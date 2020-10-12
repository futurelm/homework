package main

import (
	"log"
	"math"
	"strconv"

	//"flag"
	"flag"
)

func main() {
	var file = flag.String("file", "input.txt", "path to the input file")
	var subFileCount = flag.Int("count", 10, "number of sub files")
	flag.Parse()
	var nonRepeatedIndex int64 = math.MaxInt64
	var nonRepeatedString = ""
	files := splitHugeFileToN(*file, *subFileCount)
	// for-range files and find the non-repeated string
	for i, f := range files {
		for k, v := range readFileToWordMap(f.Name()) {
			if v.count == 1 {
				ind, _ := strconv.ParseInt(v.index, 10, 64)
				if ind < nonRepeatedIndex {
					nonRepeatedIndex = ind
					nonRepeatedString = k
				}
			}
		}
		files[i].Close()
	}
	log.Printf("first non repeated str is %s, index is %d", nonRepeatedString, nonRepeatedIndex)
}
