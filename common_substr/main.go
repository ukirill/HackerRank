package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	letters := map[rune]interface{}{}
	for _, l := range s1 {
		if _, ok := letters[l]; !ok {
			letters[l] = nil
		}
		if len(letters) == 26 {
			break
		}
	}

	for _, l := range s2 {
		if _, ok := letters[l]; ok {
			return "YES"
		}
	}
	return "NO"
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
		s1 := readLine(reader)

		s2 := readLine(reader)

		result := twoStrings(s1, s2)

		fmt.Fprintf(writer, "%s\n", result)
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
