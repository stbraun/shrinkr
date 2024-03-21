/*
Copyright © 2024 Stefan Braun sb@action.ms

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package util

import (
	"fmt"
	"os"
)

type Stats struct {
	count int
	iSize int64
	oSize int64
}

func NewStats() *Stats {
	return &Stats{
		count: 0,
		iSize: 0,
		oSize: 0,
	}
}

func (s *Stats) AddSizes(isize, osize int64) {
	s.iSize += isize
	s.oSize += osize
	s.count++
}

// Calculates the saved space.
func (s *Stats) SizeReducedBy() int64 {
	return s.iSize - s.oSize
}

// Returns the number of processed files.
func (s *Stats) Count() int {
	return s.count
}

// Returns the cumulated sizes of original files.
func (s *Stats) CumulatedSizesOfOriginalFiles() int64 {
	return s.iSize
}

// Returns the cumulated sizes of the shrinked files.
func (s *Stats) CumulatedSizesOfShrinkedFiles() int64 {
	return s.oSize
}

func GetFileSize(name string) int64 {
	fstat, err := os.Stat(name)
	if err != nil {
		panic(err)
	}
	return fstat.Size()
}

func FormatFileSize(size int64) string {
	if size > 1024 {
		sz := float64(size) / 1024
		if sz > 1024 {
			sz = sz / 1024
			return fmt.Sprintf("%.2f MB", sz)
		}
		return fmt.Sprintf("%.2f KB", sz)
	}
	return fmt.Sprintf("%d B", size)
}
