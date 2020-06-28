package regdis

import (
	"strconv"
)

type Element struct {
	BitOffset  uint8
	BitSize    uint8
	ResetValue uint64
	Name       string
	Type       string
	Desc       string
	Semantic   map[uint64]string
}

var SemanticEnable = map[uint64]string{
	0: "Disabled",
	1: "Enabled",
}

func Dissect(reg uint64, elements []Element) (ret string) {
	var totalBitSize uint8
	for _, e := range elements {
		if uint8(e.BitOffset+e.BitSize) > totalBitSize {
			totalBitSize = uint8(e.BitOffset + e.BitSize)
		}
	}
	for _, e := range elements {
		var mask uint64
		var i uint8
		for i = 0; i < e.BitSize; i++ {
			mask += 1 << i
		}
		val := uint64(reg & (mask << e.BitOffset))
		ret += displayBits(totalBitSize, e.BitOffset, e.BitSize, val)
		ret += e.Name
		if e.Desc != "" {
			ret += " : " + e.Desc
		}
		val = val >> e.BitOffset
		if e.Semantic[val] != "" {
			ret += " : " + e.Semantic[val]
		}
		ret += "\n"
	}
	return
}

func displayBits(tot, off, size uint8, val uint64) (ret string) {
	var i uint8
	for i = tot; i > 0; i-- {
		if ((i - 1) < off) || ((i - 1) >= (off + size)) {
			ret += "."
		} else {
			ret += strconv.Itoa(int((val >> (i - 1)) & 1))
		}
	}
	ret += " "
	return
}
