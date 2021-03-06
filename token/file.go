// Copyright (c) 2014, Rob Thornton
// All rights reserved.
// This source code is governed by a Simplied BSD-License. Please see the
// LICENSE included in this distribution for a copy of the full license
// or, if one is not included, you may also find a copy at
// http://opensource.org/licenses/BSD-2-Clause

package token

type File struct {
	base     int
	name     string
	newlines []Pos
	size     int
}

func NewFile(name string, base, size int) *File {
	return &File{
		base:     base,
		name:     name,
		newlines: make([]Pos, 0, 16),
		size:     size,
	}
}

func (f *File) AddLine(p Pos) {
	base := Pos(1)
	if p.Valid() && p >= base && p < base+Pos(f.Size()) {
		f.newlines = append(f.newlines, p)
	}
}

func (f *File) Base() int {
	return f.base
}

func (f *File) Pos(offset int) Pos {
	if offset < 0 || offset >= f.size {
		panic("illegal file offset")
	}
	return Pos(f.base + offset)
}

func (f *File) Position(p Pos) Position {
	start := Pos(0)
	col, row := int(p), 1

	for i, nl := range f.newlines {
		if p <= nl {
			col, row = int(p-start), i+1
			break
		}
		start = nl
	}

	return Position{Filename: f.name, Col: col, Row: row}
}

func (f *File) Size() int {
	return f.size
}
