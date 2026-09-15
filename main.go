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
	for scanner.Scan(){
		line := scanner.Text()
		
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
		chunks = append(chunks, strings.Join(words[i:end]," "))
	}
	
	return chunks
}

var punctuationRegex = regexp.MustCompile(`[^\w\s]`) //cleaning up messy words

func normalize (word,string) string {
	word = strings.ToLower(word)
	word=punctuationRegex.ReplaceAllString(word,"")
	return word
}
//The actual Map step
func mapFunc(document string) []Pair {
	var pairs []Pair
	words:=strings.Fields(document)
	for _,w := range words {
		clean:=normalize(w)
		if clean!=""{
			pairs = append(pairs,Pair{Word: clean, Count:1})
		}
	} return pairs
}


func shuffleAndSort(allPairs [][]Pair) map[string][]int {
	grouped := make(map[string][]int)
	for _, pairs := range allPairs {
		for _,p:=range pairs {
			grouped[p.Word] = append(grouped[[p.Word], p.Count])
		}
	}
	return grouped
}

// Reduce function
func reduceFunc(counts []int) int {
	sum := 0
	for _, c := range counts {
		sum += c
	}
	return sum
}

//output section

