package field

import (
	"metalim/advent/2017/lib/debug"
	"strings"
)

type slice1d []Cell
type slice2d []slice1d

// Slice segmented to 4 sectors from the center.
type Slice struct {
	field2d
	origin         Pos
	tl, tr, bl, br slice2d
}

// FillFromString from start position.
func (f *Slice) FillFromString(start Pos, s string) {
	for y, l := range strings.Split(s, "\n") {
		for x, r := range l {
			f.Set(start.Add(Pos{x, y}), Cell(r))
		}
	}
}

// Get cell.
func (f *Slice) Get(p Pos) Cell {
	x, y, sec := f.getSegmentPos(p)
	if y < len(*sec) && x < len((*sec)[y]) {
		return (*sec)[y][x]
	}
	return f.def
}

// Set cell.
func (f *Slice) Set(p Pos, c Cell) {
	if f.b.Empty() {
		f.origin = p
	}
	f.b = f.b.Union(Rect{p, p.Add(Pos{1, 1})})
	x, y, sec := f.getSegmentPos(p)

	if y >= len(*sec) {
		grow := cap(*sec)*2 + 16
		if y >= grow {
			grow = y + 16
		}
		debug.Trace("growing 2d", len(*sec), "to", grow, "and", cap(*sec), "to", grow)
		sec2 := make(slice2d, grow)
		copy(sec2, *sec)
		*sec = sec2
	}

	if x >= len((*sec)[y]) {
		grow := cap((*sec)[y])*2 + 16
		if x >= grow {
			grow = x + 16
		}
		debug.Trace("growing 1d", len((*sec)[y]), "to", grow, "and", cap((*sec)[y]), "to", grow)
		sec2 := make(slice1d, grow)
		copy(sec2, (*sec)[y])
		(*sec)[y] = sec2
	}
	(*sec)[y][x] = c
}

func (f *Slice) getSegmentPos(p Pos) (x, y int, sector *slice2d) {
	p = p.Sub(f.origin)
	y = p.Y
	x = p.X
	if y >= 0 {
		if x >= 0 {
			return x, y, &f.br
		}
		return -x - 1, y, &f.bl
	}
	if x >= 0 {
		return x, -y - 1, &f.tr
	}
	return -x - 1, -y - 1, &f.tl
}