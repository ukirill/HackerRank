package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the hackerlandRadioTransmitters function below.
func hackerlandRadioTransmitters(x []int32, k int32) int32 {
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	start := x[0]
	trm := x[0]
	var count int32 = 1
	for _, current := range x[1:] {
		if current <= start+k {
			trm = current
			continue
		}

		if current <= trm+k {
			continue
		}

		trm = current
		start = current
		count++
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	xTemp := strings.Split(readLine(reader), " ")

	var x []int32

	for i := 0; i < int(n); i++ {
		xItemTemp, err := strconv.ParseInt(xTemp[i], 10, 64)
		checkError(err)
		xItem := int32(xItemTemp)
		x = append(x, xItem)
	}

	result := hackerlandRadioTransmitters(x, k)

	fmt.Fprintf(writer, "%d\n", result)

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
