package pocketlog

// Level represents an available logging level.
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used
	// for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information
	// deemed valuable.
	LevelInfo
	// LevelError represents the highest logging level, only to be used to
	// to trace errors.
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "D"
	case LevelInfo:
		return "I"
	case LevelError:
		return "E"
	default:
		// Should never happen
		return ""
	}
}
