package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

var Version = "v0.1.1"

type Change struct {
	Added   int
	Removed int
}

func main() {
	numstatParser := regexp.MustCompile(`^(\d+)\s+(\d+)\s+(.+)$`)
	typesList := []string{"_mock_test.go", "_test.go", ".go", ""}
	types := make(map[string]*Change)
	for _, t := range typesList {
		types[t] = &Change{}

	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		match := numstatParser.FindStringSubmatch(scanner.Text())
		addedStr, removedStr, filename := match[1], match[2], match[3]
		added, _ := strconv.Atoi(addedStr)
		removed, _ := strconv.Atoi(removedStr)
		for _, t := range typesList {
			if strings.HasSuffix(filename, t) {
				types[t].Added += added
				types[t].Removed += removed
				break
			}
		}
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Type", "Added", "Removed"})
	for _, t := range typesList {
		tw.AppendRow(table.Row{t, types[t].Added, types[t].Removed})
	}
	fmt.Println(tw.Render())
}
