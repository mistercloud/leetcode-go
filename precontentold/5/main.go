package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	Solution(read, writer)
}

type Category struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent"`
}

type CategoryTree struct {
	Id   int            `json:"id"`
	Name string         `json:"name"`
	Next []CategoryTree `json:"next"`
}

func Solution(reader *bufio.Reader, writer *bufio.Writer) {
	defer writer.Flush()

	var total int
	fmt.Fscanln(reader, &total)
	root := make([]CategoryTree, total)
	i := 0
	for total > 0 {
		var lines int
		fmt.Fscanln(reader, &lines)
		var jsonStr string
		for lines > 0 {
			line, _ := reader.ReadString('\n')
			jsonStr += line
			lines--
		}

		var categories []Category
		json.Unmarshal([]byte(jsonStr), &categories)
		root[i] = getRoot(categories)
		total--
		i++
	}

	fmt.Fprintln(writer, "[")
	for i, item := range root {
		json, _ := json.Marshal(item)
		fmt.Fprintln(writer, string(json))
		if i < len(root)-1 {
			fmt.Fprintln(writer, ",")
		}
	}
	fmt.Fprintln(writer, "]")
}

func getRoot(categories []Category) CategoryTree {
	for _, category := range categories {
		if category.Id == 0 {
			return CategoryTree{
				category.Id,
				category.Name,
				getNext(categories, 0),
			}
		}
	}

	return CategoryTree{}
}

func getNext(categories []Category, parentId int) []CategoryTree {
	next := []CategoryTree{}
	for i := 0; i < len(categories); i++ {
		if categories[i].ParentId == parentId && categories[i].Id != 0 {
			next = append(next, CategoryTree{
				categories[i].Id,
				categories[i].Name,
				getNext(categories, categories[i].Id),
			})
		}
	}

	return next
}
