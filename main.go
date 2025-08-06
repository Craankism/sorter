package main

import (
	"fmt"
	"os"
	"bufio"
	"flag"
	"sorter/sorter"
)

func main() {
	// Define a flag for the input file
	order := flag.String("order", "asc", "Sort order: asc or desc")
	flag.Parse()
	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	// Read names from the file
	var names []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// Sort the names based on the specified order
	if *order == "asc" {
		sorter.BubbleSort(names)
	} else if *order == "desc" {
		sorter.BubbleSortDesc(names)
	} else {
		fmt.Println("Invalid sort order. Use 'asc' or 'desc'.")
		return
	}
	// Write the sorted names to the output file
	outputfile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputfile.Close()

	fmt.Fprintln(outputfile)
	for _, name := range names {
		fmt.Fprintln(outputfile, name)
	}
}