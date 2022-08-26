package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)


var ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
var ErrUnsupportedFile = errors.New("unsupported file")

func Copy(fromPath, toPath string, offset, limit int64) error {
	srcFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	dstFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	srcFileStat, err := srcFile.Stat()
	if err != nil {
		return err
	}
	srcFileSize := srcFileStat.Size()

	if limit <= 0 {
		limit = srcFileSize
	} else if limit > srcFileSize {
		limit = srcFileSize
	}
	s := io.NewSectionReader(srcFile, offset, limit)

	bar := pb.StartNew(int(limit))
	count := 1000
	buf := make([]byte, count)
	offs := 0
	for offs < int(limit) {
		read, err := s.Read(buf)
		if err == io.EOF {
			if _, err := dstFile.Write(buf[0:read]); err != nil {
				return err
			}
			for i := 0; i < read; i++ {
				bar.Increment()
			}
			bar.Finish()
			return nil
		} else if err != nil {
			return err
		}
		offs += read
		for i := 0; i < read; i++ {
			bar.Increment()
		}
		if len(buf) > read {
			_, err = dstFile.Write(buf[0:read])
			if err != nil {
				return err
			}
			bar.Finish()
			return nil
		}
		_, err = dstFile.Write(buf)
		if err != nil {
			return err
		}
	}

	bar.Finish()

	// // Place your code here.
	// dstFile, err := os.Create(toPath)
	// if err != nil {
	// 	return err
	// }
	// defer dstFile.Close()

	// srcFile, err := os.Open(fromPath)
	// if err != nil {
	// 	return ErrUnsupportedFile
	// }
	// srcFileStat, err := srcFile.Stat()
	// if err != nil {
	// 	return err
	// }
	// srcFileSize := srcFileStat.Size()
	// if offset > srcFileSize {
	// 	return ErrOffsetExceedsFileSize
	// }
	// var count  int64 = 100
	
	

	
	// if offset > 0 {
	// 	// bSrc = bSrc[offset:]
	// }
	// if limit > 0 {
	// 	if srcFileSize < limit {
	// 		limit = srcFileSize
	// 	}
	// 	reader := io.LimitReader(srcFile, limit)
	// 	readerB := io.LimitReader(reader, count)
	// 	bar := pb.Full.Start64(count)
	// 	barReader := bar.NewProxyReader(readerB)
	// 	io.Copy(dstFile, barReader)
	// 	bar.Finish()

	// 	return nil
	// }
	// // file.Write(bSrc)
	// // reader := io.LimitReader(srcFile, count)
	// // bar := pb.Full.Start64(count)
	// // barReader := bar.NewProxyReader(reader)
	// // io.Copy(dstFile, barReader)
	// // bar.Finish()
	return nil
}
