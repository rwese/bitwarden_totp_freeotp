package main

import (
	"bitwarden_totp_freeotp/convert"
	"flag"
	"fmt"
	"os"
)

func main() {
	inputFile := flag.String("input", "", "input JSON file (mandatory)")
	outputFile := flag.String("output", "", "output text file (mandatory)")

	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	input, err := os.Open(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open input file: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	output, err := os.Create(*outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create output file: %v\n", err)
		os.Exit(1)
	}
	defer output.Close()

	outputData, err := convert.Convert(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not convert data: %v\n", err)
		os.Exit(1)
	}

	_, err = output.Write(outputData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not write to output file: %v\n", err)
		os.Exit(1)
	}
}
