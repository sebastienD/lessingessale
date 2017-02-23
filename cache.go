package main

func (c*Cache) Insert (v Video) {
	c.Capacity -=  v.Size
}