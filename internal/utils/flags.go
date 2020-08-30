package utils

import (
	"fmt"
	"strings"
)

// AppendToFlag is a quick function that appends content to a flag if is already defined
func AppendToFlag(flags []string, flag, content string) []string {
	// no content to append to flag
	if content == "" {
		return flags
	}

	var match bool
	for index, flagEntry := range flags {
		if strings.Contains(flagEntry, flag) {
			flagEntry += " " + content
			flags[index] = flagEntry
			match = true
			break
		}
	}
	// Flag not found, so appends new flag to list
	if !match {
		flags = append(flags, fmt.Sprintf("%s=%s", flag, content))
	}

	return flags
}

// TODO func SetFlag

// TODO MergeFlags
