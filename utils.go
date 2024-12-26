package main

import (
	"fmt"
	"strconv"
)

func hexColorToEscapeSequence(hex string) (string, error) {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) != 6 {
		return "", fmt.Errorf("'%s' is not a valid hex color", hex)
	}

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return "", err
	}

	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return "", err
	}

	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b), nil
}
