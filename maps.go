package main

import "fmt"

var unit_map = map[string]int{}

func InitUnitMap() {
	unit_map["quarter_dozen"] = 3
	unit_map["half_dozen"] = 6
	unit_map["dozen"] = 12
	unit_map["small_gross"] = 120
	unit_map["gross"] = 144
	unit_map["great_gross"] = 1728
}

func GetUnitMap() map[string]int {
	return unit_map
}

func NewBill() map[string]int {
	return make(map[string]int)
}

func AddMapItem(bill, units map[string]int, item_name, quantity_string string) bool {
	if units[quantity_string] != 0 {
		if bill[item_name] != 0 {
			var old_quantity = bill[item_name]
			bill[item_name] += units[quantity_string]
			fmt.Println(fmt.Sprintf("increased '%s' quantity from %d to %d", item_name, old_quantity, bill[item_name]))
		} else {
			bill[item_name] = units[quantity_string]
			fmt.Println(fmt.Sprintf("added new item '%s' with quantity %d", item_name, bill[item_name]))
		}
		return true
	}
	return false
}

func RemoveMapItem(bill, units map[string]int, item_name, quantity_string string) bool {
	// Quantity DNE
	if units[quantity_string] == 0 {
		return false
	}
	if bill[item_name] != 0 {
		var old_quantity = bill[item_name]
		var new_quantity = bill[item_name] - units[quantity_string]
		// Invalid quantity after removal
		if new_quantity < 0 {
			fmt.Println(fmt.Sprintf("invalid quantity for '%s'; current quantity is %d and removing %d would leave %d", item_name, old_quantity, units[quantity_string], new_quantity))
			return false
		} else if new_quantity == 0 {
			delete(bill, item_name)
			fmt.Println(fmt.Sprintf("removed item '%s' from bill", item_name))
		} else {
			bill[item_name] = new_quantity
			fmt.Println(fmt.Sprintf("decreased '%s' quantity from %d to %d", item_name, old_quantity, bill[item_name]))
		}
		return true
	}
	fmt.Println(item_name, "not found, nothing to remove")
	return false
}

func GetMapItem(bill map[string]int, item_name string) (int, bool) {
	if bill[item_name] != 0 {
		return bill[item_name], true
	}
	return 0, false
}

func MapsTest() {
	InitUnitMap()
	units := GetUnitMap()
	fmt.Println("Unit map:", units)
	bill := NewBill()
	fmt.Println("New (empty) bill:", bill)
	AddMapItem(bill, units, "carrot", "dozen")
	AddMapItem(bill, units, "carrot", "dozen")
	AddMapItem(bill, units, "celery", "quarter_dozen")
	RemoveMapItem(bill, units, "carrot", "dozen")
	RemoveMapItem(bill, units, "celery", "dozen")
	qty, ok := GetMapItem(bill, "celery")
	fmt.Printf("there is %d celery in the bill, and it returned %t\n", qty, ok)
}
