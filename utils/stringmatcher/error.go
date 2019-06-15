package stringmatcher

import "errors"

// ErrNoMatch means there was no pattern matched to given data
var ErrNoMatch = errors.New("No pattern matched")
