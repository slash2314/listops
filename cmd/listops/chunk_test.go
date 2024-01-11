package main

import (
	"reflect"
	"testing"
)

type ChunkTest struct {
	delimiter string
	input     string
	output    [][]string
	size      int
}

var chunkTests = []ChunkTest{
	{delimiter: ",", input: "1,2,3", output: [][]string{{"1"}, {"2"}, {"3"}}, size: 1},
	{delimiter: ",", input: "1,2,3,4,5,6", output: [][]string{{"1", "2"}, {"3", "4"}, {"5", "6"}}, size: 2},
	{delimiter: ",", input: "1,2,3,4,5,6,7", output: [][]string{{"1", "2"}, {"3", "4"}, {"5", "6"}, {"7"}}, size: 2},
	{delimiter: "\n", input: "1\n2\n3\n4\n5\n6", output: [][]string{{"1", "2"}, {"3", "4"}, {"5", "6"}}, size: 2},
}

func TestChunk(t *testing.T) {
	for _, ct := range chunkTests {
		testOutput := Chunk(ct.input, ct.delimiter, ct.size)

		if !reflect.DeepEqual(ct.output, testOutput) {
			t.Fatalf("Expected: %s Got: %s\n", ct.output, testOutput)
		}
	}
}
