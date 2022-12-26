package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen":           12,
		"small_gross":     120,
		"gross":           144,
		"great_gross":     1728}

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unitName string) bool {
	unitValue, unitExists := units[unitName]
	if unitExists {
		bill[item] += unitValue
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill
/*
4. Remove an item from the customer bill
To implement this, you'll need to:

Return false if the given item is not in the bill
Return false if the given unit is not in the units map.
Return false if the new quantity would be less than 0.
If the new quantity is 0, completely remove the item from the bill then return true.
Otherwise, reduce the quantity of the item and return true.
bill := NewBill()
units := Units()
ok := RemoveItem(bill, units, "carrot", "dozen")
fmt.Println(ok)
*/
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemBillValue, itemInTheBillExists := bill[item]
	unitValue, unitNameExists := units[unit]

	if !itemInTheBillExists || !unitNameExists || itemBillValue-unitValue < 0 {
		return false
	}

	delete(bill, item)
	if bill[item] == 0 {
		return true
	} else {
		return false
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
/*
5. Return the quantity of a specific item that is in the customer bill
To implement this, you'll need to:

Return 0 and false if the item is not in the bill.
Otherwise, return the quantity of the item in the bill and true.
bill := map[string]int{"carrot": 12, "grapes": 3}
qty, ok := GetItem(bill, "carrot")
fmt.Println(qty)
// Output: 12
fmt.Println(ok)
// Output: true
*/
func GetItem(bill map[string]int, item string) (int, bool) {
	itemInTheBill, itemExistsInTheBill := bill[item]
	if !itemExistsInTheBill {
		return 0, false
	} else {
		return itemInTheBill, true
	}
}
