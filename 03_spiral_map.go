package main

import (
	"fmt"
	"metalim/advent/2017/lib/debug"
	"metalim/advent/2017/lib/field"
	"metalim/advent/2017/lib/source"
	"time"

	. "github.com/logrusorgru/aurora"
)

type map2d = field.Slice

// note: start from 0
/*
16 15 14 13 12
17  4  3  2 11
18  5  0  1 10
19  6  7  8  9
20 21 22 23 24
*/
func walkSpiral(f field.Field, fn func(int, field.Pos) bool) field.Pos {
	var d field.Dir4 // turn in this direction if possible.
	p := field.Pos{0, 0}
	for i := 0; fn(i, p); i++ {
		if f.Get(field.Step4(p, d)) == 0 {
			d = (d + 3) & 3 // turn left.
		}
		p = field.Step(p, (d+1)%4)
	}
	return p
}

func main() {
	var ins source.Inputs

	ins = ins.Test(1, `23`, `2`)
	ins = ins.Test(1, `1024`, `31`)
	ins = ins.Test(2, `60`, `122`)

	for in := range ins.Advent(2017, 3) {
		fmt.Println(Brown("\n" + in.Name))
		n := in.Ints()[0]
		fmt.Println(Black(n).Bold())

		if in.Part(1) {
			f := map2d{}
			p := walkSpiral(&f, func(i int, p field.Pos) bool {
				debug.LogT(time.Second, i, p)
				f.Set(p, i+1)
				return i+1 < n
			})
			in.SubmitInt(1, field.Manh(p, field.Pos{}))
		}

		if in.Part(2) {
			f := map2d{}
			p := walkSpiral(&f, func(i int, p field.Pos) bool {
				debug.LogT(time.Second, i, p)
				if i == 0 {
					f.Set(p, 1)
					return true
				}
				sum := 0
				for d := field.Dir8(0); d < 8; d++ {
					sum += f.Get(field.Step8(p, d))
				}
				f.Set(p, sum)
				return sum <= n
			})
			in.SubmitInt(2, f.Get(p))
		}
	}
}
