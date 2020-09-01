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

	var (
		extractedFlag     string
		match, simpleFlag bool
	)
	for index, flagEntry := range flags {
		// checks if is simple flag or has an assignation
		if asigIndex := strings.IndexRune(flagEntry, '='); asigIndex != -1 {
			// extracts flag
			extractedFlag = flagEntry[:asigIndex]
			simpleFlag = false
		} else {
			extractedFlag = flagEntry
			simpleFlag = true
		}

		// if flags are equal appends content to existing assignation
		if extractedFlag == flag {
			if !simpleFlag {
				flagEntry += " " + content
			} else {
				flagEntry = fmt.Sprintf("%s=%s", flag, content)
			}
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
