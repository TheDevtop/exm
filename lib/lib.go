package lib

import "io"

// Shared messages
const (
	ErrCantStart = "Error: %s, can't start\n"
	ErrObject    = "Error: %s, for object %s\n"
	ErrNoConfig  = "Error: Config file not found"
	InfoSignal   = "Info: Caught signal (%s)\n"
)

// Runner functions
type Runner func(SourceEntry, func(string, io.ReadCloser))
