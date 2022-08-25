package main

import (
	"errors"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

// ErrUnsupportedFile = errors.New("unsupported file").

var ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.
	file, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer file.Close()
	bSrc, _ := os.ReadFile(fromPath)
	if int(offset) > len(bSrc) {
		return ErrOffsetExceedsFileSize
	}
	count := 100

	// create and start new bar
	bar := pb.StartNew(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}

	// finish bar
	bar.Finish()
	if offset > 0 {
		bSrc = bSrc[offset:]
	}
	if limit > 0 {
		if len(bSrc) < int(limit) {
			limit = int64(len(bSrc))
		}
		bSrc = bSrc[:limit]
	}
	file.Write(bSrc)
	return nil
}
