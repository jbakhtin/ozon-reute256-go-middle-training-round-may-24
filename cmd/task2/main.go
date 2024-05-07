package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, start, end int
	var str, stc string
	var err error

	_, err = fmt.Fscan(in, &str, &n)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(in, &start, &end, &stc)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		builder := &strings.Builder{}

		builder.WriteString(str[:start-1])
		builder.WriteString(stc)
		builder.WriteString(str[end:])

		str = builder.String()
	}

	_, err = fmt.Fprintln(out, str)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
