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
	"reflect"
	"testing"
)

func TestNewStats(t *testing.T) {
	tests := []struct {
		name string
		want *Stats
	}{
		{"initialize", &Stats{count: 0, iSize: 0, oSize: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStats_AddSizes(t *testing.T) {
	type fields struct {
		count int
		iSize int64
		oSize int64
	}
	type args struct {
		isize int64
		osize int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stats
	}{
		{"initial", fields{count: 0, iSize: 0, oSize: 0}, args{isize: 2000, osize: 500}, &Stats{count: 1, iSize: 2000, oSize: 500}},
		{"new files", fields{count: 1, iSize: 2000, oSize: 500}, args{isize: 1200, osize: 400}, &Stats{count: 2, iSize: 3200, oSize: 900}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stats{
				count: tt.fields.count,
				iSize: tt.fields.iSize,
				oSize: tt.fields.oSize,
			}
			s.AddSizes(tt.args.isize, tt.args.osize)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("NewStats() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestStats_SizeReducedBy(t *testing.T) {
	type fields struct {
		count int
		iSize int64
		oSize int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{"initial", fields{count: 0, iSize: 0, oSize: 0}, 0},
		{"file added", fields{count: 2, iSize: 3200, oSize: 900}, 2300},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stats{
				count: tt.fields.count,
				iSize: tt.fields.iSize,
				oSize: tt.fields.oSize,
			}
			if got := s.SizeReducedBy(); got != tt.want {
				t.Errorf("Stats.SizeReducedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatFileSize(t *testing.T) {
	type args struct {
		size int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Bytes", args{size: 750}, "750 B"},
		{"KBytes", args{size: 5000}, "4.88 KB"},
		{"MBytes", args{size: 5000000}, "4.77 MB"},
		{"GBytes", args{size: 5000000000}, "4768.37 MB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatFileSize(tt.args.size); got != tt.want {
				t.Errorf("formatFileSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
