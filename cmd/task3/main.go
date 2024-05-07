package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	requestTypeGetMessage = iota + 1
	requestTypeSetMessage
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	in.Split(bufio.ScanWords)
	in.Scan()

	var n, q int // n - users count, q - requests count
	var err error

	n, err = strconv.Atoi(in.Text())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	in.Scan()

	q, _ = strconv.Atoi(in.Text())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	messages := make(map[int]int, n/2)
	personalMsgID := 1
	systemMsgID := 0

	var requestType, userID int

	for i := 0; i < q; i++ {
		in.Scan()
		requestType, err = strconv.Atoi(in.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		in.Scan()
		userID, err = strconv.Atoi(in.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch requestType {
		case requestTypeGetMessage:
			if val, ok := messages[userID]; ok {
				if val > systemMsgID {
					_, err = out.WriteString(strconv.Itoa(val) + "\n")
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				} else {
					_, err = out.WriteString(strconv.Itoa(systemMsgID) + "\n")
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
			} else {
				if systemMsgID > 0 {
					_, err = out.WriteString(strconv.Itoa(systemMsgID) + "\n")
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				} else {
					_, err = out.WriteString("0\n")
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
			}
		case requestTypeSetMessage:
			if userID == 0 {
				systemMsgID = personalMsgID
			} else {
				messages[userID] = personalMsgID
			}

			personalMsgID++
		}
	}
}
