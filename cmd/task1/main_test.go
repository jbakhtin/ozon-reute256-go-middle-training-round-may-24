package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFoo(t *testing.T) {
	testFiles := []string{"1"}
	for _, file := range testFiles {
		t.Run(file, func(t *testing.T) {
			// Файл с входными данными
			inputFileName := fmt.Sprintf("%s", file)
			inputFile, err := os.Open(inputFileName)
			if err != nil {
				t.Fatalf("Error opening file %s: %v", inputFileName, err)
			}
			defer inputFile.Close()

			// Файл с выходными данными для сравнения результата работы программы
			outputToEqualFileName := fmt.Sprintf("%s.%s", file, "a")
			outputToEqualFile, err := os.Open(outputToEqualFileName)
			if err != nil {
				t.Fatalf("Error opening file %s: %v", outputToEqualFileName, err)
			}
			defer outputToEqualFile.Close()

			outputToEqualFileReader := bufio.NewReader(outputToEqualFile)

			os.Stdin = inputFile

			// Создаем канал (pipe)
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			// Запускаем main функцию
			main()
			// Закрываем конец записи канала
			w.Close()
			// Считываем данные из конца чтения канала
			var buf bytes.Buffer
			io.Copy(&buf, r)
			// Восстанавливаем stdout
			os.Stdout = oldStdout

			var outputToEqualBuf bytes.Buffer
			_, err = outputToEqualBuf.ReadFrom(outputToEqualFileReader)

			if !bytes.Equal(outputToEqualBuf.Bytes(), buf.Bytes()) {
				t.Error("error", outputToEqualBuf.Bytes(), buf.Bytes())
			}
		})
	}
}
