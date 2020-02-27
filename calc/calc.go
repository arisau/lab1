// Provides console calculator implemention
package calc

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)


type Calculator struct{
	Config Configuration
}

// Configuration for calculator
type Configuration struct{
	WelcomeMessage string
}

// Print welcome message to the console
func (c Calculator) PrintWelcome(){
	fmt.Println(c.Config.WelcomeMessage)
}

// Sums a and b
func (c Calculator) Sum(a, b float64) float64{
	return a+b
}

// Multiplexes a by b
func (c Calculator)  Multiply(a,b float64) float64{
	return a*b
}

// Divides a by b
func (c Calculator) Divide(a,b float64) (float64, error) {
	if b == 0{
		return -1 , errors.New("cannot divide by 0")
	}
	return a/b, nil
}

func (c Calculator) Pow(a,b float64) float64{
	return math.Pow(a,b)
}

// Starts calculator instance
func (c Calculator) Start(){

	fmt.Println("Welcome to console calculator")
	for true {
		var modeStr string
		var a, b float64
		fmt.Println("\n 0 - exit \n 1 - sum \n 2 - subtract\n 3 - multiply \n 5 - divide")

		fmt.Scan(&modeStr)

		mode, err := strconv.Atoi(modeStr)
		if err != nil {
			fmt.Println("Wrong selection.")
			continue
		}

		if mode == 0{
			os.Exit(0)
		}

		fmt.Print("A: ")
		fmt.Scan(&a)
		fmt.Print("\nB: ")
		fmt.Scan(&b)

		switch mode {
		case 1:
			fmt.Println(c.Sum(a, b))
			break
		case 2:
			fmt.Println(c.Sum(a, -b))
			break
		case 3:
			fmt.Println(c.Multiply(a, b))
			break
		case 4:
			res, err := c.Divide(a, b)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Println(res)
			break
		case 5:
			fmt.Println(c.Pow(a,b))
			break
		default:
			fmt.Println("Wrong selection.")
		}
		fmt.Println()
	}
}

