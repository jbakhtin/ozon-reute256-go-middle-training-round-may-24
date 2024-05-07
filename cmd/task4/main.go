package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
)

//4
//20 10 20 30
//2 1 2 4

//3
//5 7 6
//1 1 1

type MySlice []int

func (a MySlice) Len() int           { return len(a) }
func (a MySlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MySlice) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	in.Split(bufio.ScanWords)

	in.Scan()
	testsCount, _ := strconv.Atoi(in.Text())

	for i := 0; i < testsCount; i++ {
		in.Scan()
		runnerCount, _ := strconv.Atoi(in.Text())

		timesTable := make([]int, runnerCount)
		sortedTimesTable := make(MySlice, runnerCount)
		places := make(map[int]int, runnerCount)

		for index, _ := range sortedTimesTable {
			in.Scan()
			sortedTimesTable[index], _ = strconv.Atoi(in.Text())
		}

		copy(timesTable, sortedTimesTable)
		sort.Sort(sortedTimesTable)

		medalNumber := 1
		k := 1
		places[sortedTimesTable[0]] = medalNumber
		for j := 1; j < runnerCount; j++ {
			if math.Abs(float64(sortedTimesTable[j]-sortedTimesTable[j-1])) > 1 {
				medalNumber += k
				k = 1
			} else {
				k++
			}

			places[sortedTimesTable[j]] = medalNumber
		}

		out.WriteString(strconv.Itoa(places[timesTable[0]]))
		for j := 1; j < runnerCount; j++ {
			out.WriteString(" " + strconv.Itoa(places[timesTable[j]]))
		}
		out.WriteString("\n")
	}
}
