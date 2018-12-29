package main

import (
	"encoding/hex"
	"fmt"
	"metalim/advent/2017/lib/source"
	"strings"

	. "github.com/logrusorgru/aurora"
)

func knotHash(s string) string {
	sn := []byte(s)
	sn = append(sn, 17, 31, 73, 47, 23)

	const dim = 256
	var list [dim]byte
	for i := range list {
		list[i] = byte(i)
	}

	var pos, skip int
	for round := 0; round < 64; round++ {
		for _, n := range sn {
			for i := 0; i < int(n/2); i++ { // reverse chunk of circle.
				a := (pos + i) % dim
				b := (pos + int(n) - 1 - i) % dim
				list[a], list[b] = list[b], list[a]
			}
			pos = (pos + int(n) + skip) % dim
			skip++
		}
	}

	// sparse -> dense
	var bytes [16]byte
	for i := range bytes {
		for _, j := range list[i*16 : i*16+16] {
			bytes[i] ^= j
		}
	}
	return hex.EncodeToString(bytes[:])
}

func main() {
	var ins source.Inputs

	ins = ins.Test(1, `3, 4, 1, 5`, `12`)
	ins = ins.Test(2, "", `a2582a3a0e66e6e86e3812dcb672a272`)
	ins = ins.Test(2, "AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd")

	for p := range ins.Advent(2017, 10) {
		fmt.Println(Brown("\n" + p.Name))
		fmt.Println(Black(p.Val).Bold())

		if p.Part(1) {
			sn := p.Ints()

			length := 256
			if strings.Contains(p.Name, "test") {
				length = 5
			}

			list := make([]int, length)
			for i := range list {
				list[i] = i
			}

			var pos, skip int
			for _, n := range sn {
				// reverse
				for i := 0; i < n/2; i++ {
					a := (pos + i) % length
					b := (pos + n - 1 - i) % length
					list[a], list[b] = list[b], list[a]
				}
				pos = (pos + n + skip) % length
				skip++
			}

			p.SubmitInt(1, list[0]*list[1])
		}

		if p.Part(2) {
			p.Submit(2, knotHash(p.Val))
		}

	}
}
