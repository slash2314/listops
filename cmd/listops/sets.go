package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slash2314/listops/internal/templates"
	"sort"
	"strings"
)

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
	resultSet := make([]string, 0)
	for v := range aSet {
		if _, ok := bSet[v]; !ok {
			resultSet = append(resultSet, v)
		}
	}
	sort.Strings(resultSet)
	return resultSet
}

func bDiffA(a, b []string) []string {
	return aDiffB(b, a)
}

func intersection(a, b []string) []string {
	aSet, bSet := createSets(a, b)
	resultSet := make([]string, 0)
	for v := range aSet {
		if _, ok := bSet[v]; ok {
			resultSet = append(resultSet, v)
		}
	}
	sort.Strings(resultSet)
	return resultSet
}

func createSets(a, b []string) (map[string]bool, map[string]bool) {
	aSet := make(map[string]bool)
	bSet := make(map[string]bool)
	for _, v := range a {
		aSet[v] = true
	}
	for _, v := range b {
		bSet[v] = true
	}
	return aSet, bSet
}

func union(a, b []string) []string {
	aSet, bSet := createSets(a, b)
	for k := range aSet {
		bSet[k] = true
	}
	resultSet := make([]string, 0)
	for v := range bSet {
		resultSet = append(resultSet, v)
	}
	sort.Strings(resultSet)
	return resultSet
}
