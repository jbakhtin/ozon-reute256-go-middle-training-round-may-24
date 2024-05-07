package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type MySlice []int

func (a MySlice) Len() int           { return len(a) }
func (a MySlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MySlice) Less(i, j int) bool { return a[i] > a[j] }

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{items: []int{}}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem
}

func (s *Stack) Peek() int {
	lastIndex := len(s.items) - 1
	return s.items[lastIndex]
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	in.Split(bufio.ScanWords)

	in.Scan()
	friendsCount, _ := strconv.Atoi(in.Text())
	in.Scan()
	cardsCount, _ := strconv.Atoi(in.Text())

	var friendCards MySlice = make([]int, friendsCount)
	var friendCardsMySliceSorted MySlice = make([]int, friendsCount)
	var cardsMap = make(map[int]*Stack, friendsCount)

	for i := 0; i < friendsCount; i++ {
		in.Scan()
		friendCards[i], _ = strconv.Atoi(in.Text())
	}

	copy(friendCardsMySliceSorted, friendCards)
	sort.Sort(friendCardsMySliceSorted)

	for i := 0; i < friendsCount; i++ {
		if cardsCount == friendCardsMySliceSorted[i] {
			out.WriteString("-1")
			return
		}

		if val, ok := cardsMap[friendCardsMySliceSorted[i]]; ok {
			val.Push(cardsCount)
			//fmt.Println("push to FIFO")

			cardsMap[friendCardsMySliceSorted[i]] = val
		} else {
			val = NewStack()
			val.Push(cardsCount)

			cardsMap[friendCardsMySliceSorted[i]] = val

			//fmt.Println("create FIFO")
			//fmt.Println("push to FIFO")
		}
		cardsCount--
	}

	item := cardsMap[friendCards[0]]
	out.WriteString(strconv.Itoa(item.Pop()))
	for i := 1; i < friendsCount; i++ {
		item := cardsMap[friendCards[i]]
		out.WriteString(" " + strconv.Itoa(item.Pop()))
	}

	//TODO: прорабоать невозможные случаи
}
