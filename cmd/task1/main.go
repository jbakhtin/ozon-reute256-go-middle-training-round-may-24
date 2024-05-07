package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, a, b int
	var err error

	_, err = fmt.Fscan(in, &t)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < t; i++ {
		_, err = fmt.Fscan(in, &a, &b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = fmt.Fprintln(out, a+b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
