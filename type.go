package main

import (
	"bufio"
	"os"
)

type word struct {
	value []byte
	index int64
}

type wordMeta struct {
	index string
	count int
}

type fileWrite struct {
	*os.File
	wr *bufio.Writer
}
