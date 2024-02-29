package cmd

import (
	"fmt"
	"io"
	"os"
)

// Copied from StackOverflow to read the last line of a file
// https://stackoverflow.com/questions/17863821/how-to-read-last-lines-from-a-big-file-with-go-every-10-secs
// Preconditions:
// filepath is a valid path to a file
func getLastLineWithSeek(filepath string) string {
	fileHandle, err := os.Open(filepath)
	if err != nil {
		panic("Cannot open file")
	}
	defer fileHandle.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	// Avoid looping if filesize is 0
	if stat.Size() <= 0 {
		return ""
	}
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	return line
}
