package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	var numberOfOccurrences int64 = 0
	if len(s) == 1 && s == "a"{
		return n
	}
	var sLength = int64(len(s))
	var a1 = n / sLength
	var a2 = int64(strings.Count(s, "a"))
	var a3 = a2 * a1
	var a4 = a1 * sLength
	var a5 = math.Abs(float64(n - a4))
	var stringSlice = strings.SplitAfter(s, "")
	var i int64
	for i = 0; i < int64(a5); i++ {
		if stringSlice[i] == "a" {
			a3++
		}
	}
	numberOfOccurrences = a3
	/*var stringSlice = strings.SplitAfter(s, "")
	var i int64
	var counter = 0

	for i = 0; i <= n; i++ {
		if stringSlice[counter] == "a" {
			numberOfOccurrences++
		}
		counter++
		if counter >= len(stringSlice) {
			counter = 0
		}
	}*/
	return numberOfOccurrences
}

func main() {

	s := "aba"
	var n int64 = 10
	result := repeatedString(s, n)
fmt.Println(result)
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
