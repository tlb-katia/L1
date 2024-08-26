package main

import "fmt"

type transport interface {
	hello()
	move()
	speed(s float32)
}

type user struct{}

func (u *user) pickTransport(tran transport, speed float32) {
	tran.hello()
	tran.move()
	tran.speed(speed)
}

type car struct{}

type bike struct{}

func (c *car) hello() {
	fmt.Println("hello, today you've picked a car")
}

func (c *car) speed(s float32) {
	fmt.Println("You're driving ", s, " km/h")
}

func (c *car) move() {
	fmt.Println("Let's go back home!")
}

type bikeAdapter struct {
	bikeAdapt *bike
}

func (b *bikeAdapter) hello() {
	fmt.Println("hello, today you've picked a bike")
}

func (b *bikeAdapter) move() {
	fmt.Println("Let's get risky!")
}

func (b *bikeAdapter) speed(s float32) {
	fmt.Println("You're driving ", s, " km/h")
}

func main() {
	u := &user{}
	car := &car{}
	bike := &bikeAdapter{bikeAdapt: &bike{}}

	u.pickTransport(car, 90)
	fmt.Println()
	u.pickTransport(bike, 180)
}
