package main

import (
	"reflect"
	"testing"
)

type SetTestData struct {
	a      []string
	b      []string
	output []string
}

var unionTestData = []SetTestData{
	{a: []string{"a", "b"}, b: []string{"c", "d"}, output: []string{"a", "b", "c", "d"}},
}
var intersectionTestData = []SetTestData{
	{a: []string{"a", "b"}, b: []string{"b", "c"}, output: []string{"b"}},
}
var diffTestData = []SetTestData{
	{a: []string{"a", "b"}, b: []string{"b", "c"}, output: []string{"a"}},
}

type CreateSetTestData struct {
	a    []string
	b    []string
	set1 map[string]bool
	set2 map[string]bool
}

var createSetsTestData = []CreateSetTestData{
	{a: []string{"a", "b"}, b: []string{"b", "c"}, set1: map[string]bool{"a": true, "b": true}, set2: map[string]bool{"b": true, "c": true}},
}

type ExtractListTestData struct {
	l, delimiter string
	output       []string
}

var extractListTestData = []ExtractListTestData{
	{"1,2,3", ",", []string{"1", "2", "3"}},
	{"1, 2, 3", ",", []string{"1", "2", "3"}},
}

func TestUnion(t *testing.T) {
	for _, st := range unionTestData {
		testOutput := union(st.a, st.b)
		if !reflect.DeepEqual(testOutput, st.output) {
			t.Fatalf("Expected: %s Got: %s\n", st.output, testOutput)
		}
	}
}

func TestIntersection(t *testing.T) {
	for _, st := range intersectionTestData {
		testOutput := intersection(st.a, st.b)
		if !reflect.DeepEqual(testOutput, st.output) {
			t.Fatalf("Expected: %s Got: %s\n", st.output, testOutput)
		}
	}
}

func TestDiff(t *testing.T) {
	for _, st := range diffTestData {
		testOutput := aDiffB(st.a, st.b)
		if !reflect.DeepEqual(testOutput, st.output) {
			t.Fatalf("Expected: %s Got: %s\n", st.output, testOutput)
		}
	}
}

func TestCreateSets(t *testing.T) {
	for _, st := range createSetsTestData {
		set1, set2 := createSets(st.a, st.b)
		if !reflect.DeepEqual(st.set1, set1) || !reflect.DeepEqual(st.set2, set2) {
			t.Fatalf("Expected: %v and %v Got: %v and %v\n", st.set1, st.set2, set1, set2)
		}
	}
}

func TestExtractList(t *testing.T) {
	for _, st := range extractListTestData {
		testOutput := extractList(st.l, st.delimiter)
		if !reflect.DeepEqual(testOutput, st.output) {
			t.Fatalf("Expected: %s Got: %s\n", st.output, testOutput)
		}
	}
}
