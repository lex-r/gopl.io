package main

import (
	"fmt"
)

type Foot float64

func (f Foot) String() string { return fmt.Sprintf("%g ft", f) }

func FToM(f Foot) Metre { return Metre(f * 0.3048) }

type Metre float64

func (m Metre) String() string { return fmt.Sprintf("%g m", m) }

func MToF(m Metre) Foot { return Foot(m / 0.3048) }
