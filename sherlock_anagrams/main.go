package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"strconv"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

var lettersMap map[rune]byte
var h = fnv.New32a()

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	prepareLettersMap()
	freqByLen := map[int]map[uint32]byte{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			f := calculateFreq(s[i:j])
			l := j - i

			if _, ok := freqByLen[l]; !ok {
				freqByLen[l] = map[uint32]byte{}
			}
			freqByLen[l][f]++
		}
	}

	c := 0
	for _, freqs := range freqByLen {
		for _, entry := range freqs {
			i := int(entry)
			c += i * (i - 1) / 2
		}
	}

	return int32(c)
}

func calculateFreq(s string) uint32 {
	freq := make([]byte, len(letters))
	for _, r := range s {
		freq[lettersMap[r]]++
	}
	return calculateHash(freq)
}

func calculateHash(freq []byte) uint32 {
	defer h.Reset()
	h.Write(freq)
	return h.Sum32()
}

func prepareLettersMap() {
	lettersMap = make(map[rune]byte, len(letters))
	for i, r := range letters {
		lettersMap[r] = byte(i)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
