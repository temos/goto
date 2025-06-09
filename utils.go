package main

import (
	"fmt"
	"strconv"
)

// hexColorToEscapeSequence converts a hexadecimal representation of a color into an ANSI escape sequence
// which, when printed to the terminal, changes the output following this sequence to this color
func hexColorToEscapeSequence(hex string) (string, error) {
	//strip the leading hash symbol, if present
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
