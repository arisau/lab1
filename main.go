package main

import "github.com/arisau/lab1/calc"

func main(){
	c:= calc.Calculator{Config:calc.Configuration{WelcomeMessage:"Welcome to calculator"}}
	c.Start()
}