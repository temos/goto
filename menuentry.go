package main

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
)

type MenuEntry struct {
	Prefix string
	Name   string
	Value  string

	//used for searching
	SearchVector string
	//updated by the searching algorithm, used to sort the search results
	Rank int
}

func MakeMenuEntries(config *Config) ([]MenuEntry, error) {
	directoriesMenuEntries, err := makeDirectoriesMenuEntries(config.Directories)
	if err != nil {
		return nil, fmt.Errorf("make directories menu entries: %w", err)
	}

	urlsMenuEntries := makeURLsMenuEntries(config.URLs)

	menuEntries := slices.Concat(directoriesMenuEntries, urlsMenuEntries)
	return menuEntries, nil
}

func makeDirectoriesMenuEntries(directories []ConfigDirectory) ([]MenuEntry, error) {
	entries := make([]MenuEntry, 0, 30)

	//for each configured directory, add all subdirectories as menu entries
	for _, directory := range directories {
		dirEntries, err := os.ReadDir(directory.Path)
		if err != nil {
			return nil, err
		}

		for _, dirEntry := range dirEntries {
			if !dirEntry.IsDir() {
				continue
			}

			if !directory.ShowHidden && strings.HasPrefix(dirEntry.Name(), ".") {
				continue
			}

			entries = append(entries, MenuEntry{
				Prefix:       directory.Prefix,
				Name:         dirEntry.Name(),
				Value:        path.Join(directory.Path, dirEntry.Name()),
				SearchVector: "directory " + directory.Prefix + " " + dirEntry.Name(),
			})
		}
	}

	return entries, nil
}

func makeURLsMenuEntries(urls []ConfigURL) []MenuEntry {
	entries := make([]MenuEntry, 0, len(urls))

	// add each URL as a menu entry
	for _, url := range urls {
		entries = append(entries, MenuEntry{
			Prefix:       url.Prefix,
			Name:         url.Name,
			Value:        url.URL,
			SearchVector: "url " + url.Prefix + " " + url.Name + " " + url.URL,
		})
	}

	return entries
}
