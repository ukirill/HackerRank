package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	runes := make(map[rune]int)

	for _, r := range a {
		runes[r]++
	}

	for _, r := range b {
		runes[r]--
	}

	count := 0
	for _, c := range runes {
		count += abs(c)
	}

	return int32(count)
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

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
