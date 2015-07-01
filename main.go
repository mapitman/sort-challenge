package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	infile, err := os.Open("./BigFile.tsv")
	defer infile.Close()
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sort.Strings(lines)

	outfile, err := os.Create("./sorted.tsv")
	defer outfile.Close()

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(outfile)

	var lastline = ""
	for _, line := range lines {
		if line != lastline {
			fmt.Fprintln(w, line)
			lastline = line
		}
	}

	w.Flush()
}
