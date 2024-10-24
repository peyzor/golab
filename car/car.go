package car

// https://github.com/tmrts/go-patterns/blob/master/creational/builder.md

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type carBuilder struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func NewBuilder() Builder {
	return &carBuilder{}
}

func (cb *carBuilder) Color(c Color) Builder {
	cb.color = c
	return cb
}
func (cb *carBuilder) Wheels(w Wheels) Builder {
	cb.wheels = w
	return cb
}
func (cb *carBuilder) TopSpeed(s Speed) Builder {
	cb.speed = s
	return cb
}

func (cb *carBuilder) Build() Interface {
	return &car{
		color:  cb.color,
		wheels: cb.wheels,
		speed:  cb.speed,
	}
}

type car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func (c *car) Drive() error {
	fmt.Printf("Driving a %s car with %s wheels at top speed of %.2f MPH.\n", c.color, c.wheels, c.speed)
	return nil
}

func (c *car) Stop() error {
	fmt.Printf("The %s car has stopped.\n", c.color)
	return nil
}
