package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var counters = map[int32]int32{}
var freq = map[int32]int32{}

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	result := []int32{}
	for _, query := range queries {
		switch query[0] {
		case 1:
			add(query[1])
		case 2:
			del(query[1])
		case 3:
			if freq[query[1]] == 0 {
				result = append(result, 0)
			} else {
				result = append(result, 1)
			}
		}
	}
	return result
}

func add(a int32) {
	if counters[a] > 0 {
		freq[counters[a]]--
	}
	counters[a]++
	freq[counters[a]]++
}

func del(a int32) {
	if counters[a] == 0 {
		return
	}
	freq[counters[a]]--
	counters[a]--
	if counters[a] > 0 {
		freq[counters[a]]++
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
