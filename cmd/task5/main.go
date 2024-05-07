package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type Folder struct {
	//Dir     string   `json:"dir,omitempty"`
	Files   []string `json:"files,omitempty"`
	Folders []Folder `json:"folders,omitempty"`
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	in.Split(bufio.ScanLines)

	in.Scan()
	testsCount, _ := strconv.Atoi(in.Text())

	for i := 0; i < testsCount; i++ {
		in.Scan()
		rowsCount, _ := strconv.Atoi(in.Text())
		var rawJSON strings.Builder

		var stuct Folder

		for j := 0; j < rowsCount; j++ {
			in.Scan()
			rawJSON.WriteString(in.Text())
		}

		json.Unmarshal([]byte(rawJSON.String()), &stuct)

		//fmt.Println(checkFoldersRecursive(stuct))

		test := strconv.Itoa(checkFoldersRecursive(stuct, false))
		//fmt.Printf("root result: %s", test)
		out.WriteString(test + "\n")
		//fmt.Println("----")
	}
}

func checkFoldersRecursive(rootFolder Folder, all bool) int {
	var filesFound int

	if all {
		filesFound = len(rootFolder.Files)
	}

	for _, file := range rootFolder.Files {
		if strings.HasSuffix(file, ".hack") {
			filesFound = len(rootFolder.Files)
			all = true
			break
		}
	}

	for _, folder := range rootFolder.Folders {
		filesFound += checkFoldersRecursive(folder, all)
	}

	return filesFound
}
