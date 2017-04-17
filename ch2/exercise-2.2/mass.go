package main

import (
	"fmt"
)

type Pound float64

func (p Pound) String() string { return fmt.Sprintf("%g lb", p) }

func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

type Kilogram float64

func (kg Kilogram) String() string { return fmt.Sprintf("%g kg", kg) }

func KToP(kg Kilogram) Pound { return Pound(kg / 0.45359237) }
