package main

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"log"
	"os"
)

/* This program prints the zlib compression ratio for each line of stdin */

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var b bytes.Buffer
	var zwriter *zlib.Writer
	num_lines := 0
	for scanner.Scan() {
		b.Reset()
		zwriter = zlib.NewWriter(&b)
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal("error reading from input: ", err)
		}
		fmt.Fprintln(zwriter, line)
		zwriter.Close()
		full_len := len(line)
		compressed_len := b.Len()
		r := float64(full_len) / float64(compressed_len)
		fmt.Printf("%0.4f\n", r)
		num_lines++
	}
	log.Printf("Procssed %d lines\n", num_lines)
}
