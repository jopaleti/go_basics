package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type bill struct {
	name string
	items map[string]float64
	tip float64 
}

func bluePrint(n string, t float64) bill {
	b := bill{
		name: n,
		items: map[string]float64{"Cake": 1.00, "Pie": 23.78},
		tip: t,
	}
	return b
}

// RECEIVER FUNCTION --- Associated to a particular struct or data
// Format the bill
func (b bill) format() string {
	// fs --- Formatted string
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// Total
	fs += fmt.Sprintf("%-25v ...$0.2%f", "total:", total)

	return fs
}

// Function to get input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	
	name, _ := getInput("Create a new bill name: ", reader)

	b := bluePrint(name, 3000.00)
	fmt.Println("Created the bill - ", b.name)

	return b
}

// Add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// RECEIVER FUNCTION ASSOCIATED TO STRUCT BILL TO UPDATE TIP
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Prompt Options
func prompOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add Item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)

	// Switch use cases
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			prompOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		prompOptions(b)
	case "t":
		tip, _ := getInput("Enter the tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			prompOptions(b)
		}

		b.updateTip(t)

		fmt.Println("Tip added - ", tip)
		prompOptions(b)
	case "s":
		b.save()
		fmt.Println("You save the file - ", b.name)
	default: 
		fmt.Println("That was not a valid option...")
		prompOptions(b)
	}
}

// SAVE BILL
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("user_input/"+b.name+".txt", data, 0644)

	if err != nil {
		// Panic is used to stop the flow of a program and print out an error
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}

func main() {
	mybill := createBill()
	prompOptions(mybill)

	fmt.Println(mybill)
}

