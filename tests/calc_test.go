package tests

import (
	"github.com/arisau/lab1/calc"
	"math/rand"
)
import "testing"

func initEnv() calc.Calculator{
	return calc.Calculator{Config:calc.Configuration{WelcomeMessage: "Welcome to test calc"}}
}

func TestSum(t *testing.T){
	c := initEnv()

	a := rand.Float64()
	b := rand.Float64()
	sum := c.Sum(a,b)
	if sum != a+b{
		t.Errorf("Sum was incorrect, got: %f, want: %f.", sum, a+b)
	}
}

func TestSubstract(t *testing.T){
	c := initEnv()

	a := rand.Float64()
	b := rand.Float64()
	sub := c.Sum(a,-b)
	if sub != a-b{
		t.Errorf("Substraction was incorrect, got: %f, want: %f.", sub, a-b)
	}
}

func TestMult(t *testing.T){
	c := initEnv()

	a := rand.Float64()
	b := rand.Float64()
	mul := c.Multiply(a,b)
	if mul != a*b{
		t.Errorf("Multiplextion was incorrect, got: %f, want: %f.", mul, a*b)
	}
}

func TestDivide(t *testing.T){
	c := initEnv()

	a := rand.Float64()
	b := rand.Float64()
	if b == 0.0{
		b = 5.32
	}
	priv, _ := c.Divide(a,b)
	if priv != a/b{
		t.Errorf("Private was incorrect, got: %f, want: %f.", priv, a/b)
	}
}


func TestDivideZero(t *testing.T){
	c := initEnv()

	a := rand.Float64()
	b := 0.0
	priv, err := c.Divide(a,b)
	if err == nil{
		t.Errorf("Private was incorrect, expected error, got %f", priv)
	}
}