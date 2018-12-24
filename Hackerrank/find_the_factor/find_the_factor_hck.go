package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func pthFactor(n int64, p int64) int64 {
	var arr []int64
	var brr []int64
	var i, j int64
	k := p - 1
	for i = 1; i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			arr = append(arr, i)
			if n/i != i {
				brr = append(brr, n/i)
			}
		}
	}

	for i, j = 0, int64(len(brr))-1; i < j; i, j = i+1, j-1 {
		brr[i], brr[j] = brr[j], brr[i]
	}
	arr = append(arr, brr...)
	fmt.Println(arr, brr)
	if p > int64(len(arr)) {
		return 0
	} else {
		return arr[k]
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	p, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := pthFactor(n, p)

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
