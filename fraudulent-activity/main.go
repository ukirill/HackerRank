package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
	l := len(expenditure)
	td := int(d)
	if l <= td {
		return 0
	}

	freq := make([]int, 201)
	for i := 0; i < td; i++ {
		freq[expenditure[i]]++
	}

	var i1, i2 int
	if td % 2 == 0 {
		i1, i2 = td / 2 -1, td/2
	} else {
		i1, i2 = td/2, td/2
	}

	var c int32
	var m1, m2 int32
	for i := td; i < l; i++ {

		for j, k := 0, 0; k<=i1; j, k = j+1, k+freq[j]{
			m1 = int32(j)
		}
		for j, k := 0, 0; k<=i2; j, k = j+1, k+freq[j]{
			m2 = int32(j)
		}

		if expenditure[i] >= m1+m2 {
			c++
		}
		freq[expenditure[i-td]]--
		freq[expenditure[i]]++
	}

	return c
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
