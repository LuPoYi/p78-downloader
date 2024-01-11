package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 自定義比較函數，用於對含有數字的字符串進行排序
func customSort(lines []string) {
	sort.Slice(lines, func(i, j int) bool {
		partsI := strings.Split(lines[i], "_")
		partsJ := strings.Split(lines[j], "_")

		numI, errI := strconv.Atoi(partsI[0])
		numJ, errJ := strconv.Atoi(partsJ[0])

		// 如果都轉換成功，按數字比較
		if errI == nil && errJ == nil {
			return numI < numJ
		}

		// 否則，按原始字符串比較
		return lines[i] < lines[j]
	})
}

func main() {
	// 將這裡的文件路徑替換為您的文件
	inputFilePath := "./download_list_before.txt"
	outputFilePath := "./download_list_after.txt"

	// 打開輸入文件
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 讀取文件到一個切片中
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// 對切片進行排序
	customSort(lines)

	// 打開輸出文件
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// 寫入排序後的行到輸出文件
	for _, line := range lines {
		_, err := outputFile.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("文件排序完成。")
}
