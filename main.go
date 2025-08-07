package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sorter/sorter"
)

func main() {
	// Define flags
	order := flag.String("order", "asc", "Sort order: asc or desc")
	outputFileName := flag.String("output", "output.txt", "Output file name")
	debugMode := flag.Bool("debug", false, "Enable debug mode")
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
	// Create or open the output file
	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()
	
	// Sort the names based on the specified order
	switch *order {
	case "asc":
		sorter.BubbleSort(names, *debugMode)
	case "desc":
		sorter.BubbleSortDesc(names, *debugMode)
	default:
		sorter.BubbleSort(names, *debugMode)
		return
	}

	// Write the sorted names to the output file
	for _, name := range names {
		fmt.Fprintln(outputFile, name)
	}

}
