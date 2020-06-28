# go-regdis : Go Library to Dissect Register

## Overview [![GoDoc](https://godoc.org/github.com/Georges760/go-regdis?status.svg)](https://godoc.org/github.com/Georges760/go-regdis)

## Install

```
go get github.com/Georges760/go-regdis
```

## Example

```
config := []Element{}
config = append(config, Element{
	BitOffset:  1,
	BitSize:    1,
	ResetValue: 0,
	Name:       "PWR_UP",
	Type:       "R/W",
	Desc:       "",
	Semantic: map[uint64]string{
		0: "POWER DOWN",
		1: "POWER UP",
	},
})
config = append(config, Element{
	BitOffset:  0,
	BitSize:    1,
	ResetValue: 0,
	Name:       "PRIM_RX",
	Type:       "R/W",
	Desc:       "RX_TX control",
	Semantic: map[uint64]string{
		0: "PTX",
		1: "PRX",
	},
})
ret := Dissect(0x0c, config)
fmt.Println(ret)
```

## License

MIT.
