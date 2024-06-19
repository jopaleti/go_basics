package main

import (
	"fmt"
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

// RECEIVER FUNCTION ASSOCIATED TO STRUCT BILL TO UPDATE TIP
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func main() {
	myBill := bluePrint("Oluwatobi", 200.00)
	myBill.updateTip(12000.00)
	myBill.addItem("Apple", 350.00)
	fmt.Println(myBill)
	// Print the formatted string out from the receiver function above
	fmt.Println(myBill.format())

}