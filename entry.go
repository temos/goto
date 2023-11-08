package main

import (
	"os"
	"path"
)

type entry struct {
	fullPath     string
	name         string
	prefix       string
	rank         int
	searchVector string
}

func collectEntries(roots []string, prefixes []string) ([]entry, error) {
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
