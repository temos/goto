package main

import (
	"os"
	"path"
	"strings"
)

type entry struct {
	// The full path of the file
	fullPath string

	// The name of the file to show after the slash
	name string

	// The file prefix to show before the slash
	prefix string

	// A value to sort the search results by
	rank int

	// A value that combines all the values this entry should be found by
	searchVector string
}

func collectEntries(roots []string, prefixes []string, includeHidden bool) ([]entry, error) {
	if len(roots) != len(prefixes) {
		panic("prefixes must have the same length as roots")
	}

	entries := make([]entry, 0, 30)

	for rootIdx, root := range roots {
		dirEntries, err := os.ReadDir(root)
		if err != nil {
			return nil, err
		}

		for _, dirEntry := range dirEntries {
			if !dirEntry.IsDir() {
				continue
			}

			if !includeHidden && strings.HasPrefix(dirEntry.Name(), ".") {
				continue
			}

			entries = append(entries, entry{
				name:         dirEntry.Name(),
				fullPath:     path.Join(root, dirEntry.Name()),
				prefix:       prefixes[rootIdx],
				searchVector: prefixes[rootIdx] + " " + dirEntry.Name(),
			})
		}
	}

	return entries, nil
}
