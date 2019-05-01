package command

import "fmt"

type Beaker struct {
	water    float64
	salt     float64
	isMelted bool
}

func NewBeaker(water, salt float64) *Beaker {
	beaker := new(Beaker)
	beaker.water = water
	beaker.salt = salt
	beaker.Mix()
	return beaker
}

func (b *Beaker) AddWater(water float64) {
	b.water += water
}

func (b *Beaker) AddSalt(salt float64) {
	b.salt += salt
}

func (b *Beaker) Mix() {
	if (b.salt/(b.water+b.salt))*100 < 26.4 {
		b.isMelted = true
	} else {
		b.isMelted = false
	}
}

func (b *Beaker) GetWater() float64 {
	return b.water
}

func (b *Beaker) GetSalt() float64 {
	return b.salt
}

func (b *Beaker) IsMelted() bool {
	return b.isMelted
}

func (b *Beaker) Note() {
	fmt.Printf("水: %fg\n", b.water)
	fmt.Printf("食塩: %fg\n", b.salt)
	fmt.Printf("濃度: %f%%\n", (b.salt/(b.water+b.salt))*100)
}
