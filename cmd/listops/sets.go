package main

import (
	"cmp"
	"github.com/gin-gonic/gin"
	"github.com/slash2314/listops/internal/templates"
	"slices"
	"sort"
	"strings"
)

type Empty = struct{}
type Set[T cmp.Ordered] struct {
	inner map[T]Empty
}

func (t Set[T]) Insert(elems ...T) []T {
	for _, elem := range elems {
		t.inner[elem] = Empty{}
	}
	return elems
}
func (t Set[T]) Get(elem T) (T, bool) {
	_, ok := t.inner[elem]
	return elem, ok
}
func (t Set[T]) Contains(elem T) bool {
	_, ok := t.inner[elem]
	return ok
}
func (t Set[T]) Keys() []T {
	resultSet := make([]T, 0)
	for k := range t.inner {
		resultSet = append(resultSet, k)
	}
	return resultSet
}
func (t Set[T]) KeysSorted() []T {
	resultSet := make([]T, 0)
	for k := range t.inner {
		resultSet = append(resultSet, k)
	}
	slices.Sort(resultSet)
	return resultSet
}

func (t Set[T]) KeysSortedFunc(cmp func(a T, b T) int) []T {
	resultSet := make([]T, 0)
	for k := range t.inner {
		resultSet = append(resultSet, k)
	}
	slices.SortFunc(resultSet, cmp)
	return resultSet
}

func (t Set[T]) Diff(other Set[T]) Set[T] {
	resultSet := createSet[T]()
	for _, k := range t.Keys() {
		if !other.Contains(k) {
			resultSet.Insert(k)
		}
	}
	return resultSet
}

func (t Set[T]) Intersection(other Set[T]) Set[T] {
	resultSet := createSet[T]()
	for _, k := range t.Keys() {
		if other.Contains(k) {
			resultSet.Insert(k)
		}
	}
	return resultSet
}
func (t Set[T]) Union(other Set[T]) Set[T] {
	resultSet := createSet[T]()
	resultSet.Insert(t.Keys()...)
	resultSet.Insert(other.Keys()...)
	return resultSet
}

func createSet[T cmp.Ordered]() Set[T] {

	return Set[T]{inner: make(map[T]Empty)}
}

func setOperations(c *gin.Context, delimiter string, op Op) error {
	a := extractList(c.Request.PostFormValue("groupa"), delimiter)
	b := extractList(c.Request.PostFormValue("groupb"), delimiter)
	result := op(a, b)
	component := templates.SetResponder(strings.Join(result, "\n"))
	return component.Render(c.Request.Context(), c.Writer)
}

func extractList(l, delimiter string) []string {
	itemList := strings.Split(l, delimiter)
	for i := range itemList {
		itemList[i] = strings.TrimSpace(itemList[i])
	}
	sort.Strings(itemList)
	return itemList
}

func aDiffB(a, b []string) []string {
	aSet, bSet := createSets(a, b)
	return aSet.Diff(bSet).KeysSorted()
}

func bDiffA(a, b []string) []string {
	return aDiffB(b, a)
}

func intersection(a, b []string) []string {
	aSet, bSet := createSets(a, b)
	return aSet.Intersection(bSet).KeysSorted()
}

func createSets(a, b []string) (Set[string], Set[string]) {
	aSet := createSet[string]()
	bSet := createSet[string]()
	aSet.Insert(a...)
	bSet.Insert(b...)
	return aSet, bSet
}

func union(a, b []string) []string {
	aSet, bSet := createSets(a, b)
	bSet.Insert(aSet.Keys()...)
	return bSet.KeysSorted()
}
