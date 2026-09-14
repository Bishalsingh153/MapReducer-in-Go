package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type Pair struct {
	Word string
	Count int
}

func readInputFromUser() string {
	fmt.Println("Enter text (type'END' on a new line to finish):")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.scan(){
		line := scanner.Text()
		if line == "END":
		if line == "END" {
			break
		}
		lines = append(lines,line)
	}
	return strings.Join(lines," ")
}

// cutting the text into smaller pieces 
func splitIntoChunks(text string, chunkSize int) [] string {
	words := strings.Fields(text)
	var chunks []string
	for i:=0; i < len(words); i+=chunkSize{
		end := i+chunkSize
		if end > len(words) {
			end = len(words)
		}
		chunks:= append(chunks, strings.Join(words[i:end]," "))
	}
	
	return chunks
}