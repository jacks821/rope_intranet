package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
	"bufio"
	"log"
)

func GrabLines(args string) []string {
	var lines []string
	file, err := os.Open(args)
	if err != nil {
	log.Fatal(err)
		}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func prog(case_num int, va []int, vb []int) {
	fmt.Printf("Case #%d: ", case_num)
	total := 0
	for i := range va {
		for j := range vb {
			if i == j { 
				continue
			}
			if va[i] < va[j] && vb[i] > vb[j] { 
				total += 1
			}
		}
	}
	fmt.Printf("%d\n", total)
}

func main() {
	argsWithoutProgram := os.Args[1]; if argsWithoutProgram == ""{
		fmt.Println("Need a single filename argument to run this program")
	}
	b := (GrabLines(argsWithoutProgram)) 
	cases, _ := strconv.Atoi(b[0])
	index := 1
	for i := 0; i < cases; i++ {
		N, _ := strconv.Atoi(b[index])
		index += 1
		va := make([]int, N)
		vb := make([]int, N) 
		for j, line := range b[index:(index + N)] {
			lineslice := strings.Split(line, " ")
			va[j],_ = strconv.Atoi(lineslice[0])
			vb[j],_ = strconv.Atoi(lineslice[1])
		}
		prog(i+1, va, vb)
		index += N
	}
}
