package main

import "fmt"

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{name, age, address}
}

func (r Resident) HasRequiredInfo() bool {
	if r.Age < 1 || r.Age >= 130 {
		fmt.Println("invalid age")
		return false
	}
	_, ok := r.Address["street"]
	if !ok {
		fmt.Println("missing address street")
		return false
	}
	return true
}

func (r *Resident) Delete() {
	r.Address = map[string]string{}
	r.Age = 0
	r.Name = ""
}

func CountResidents(resident_list []*Resident) int {
	valid_count := 0
	total_count := 0
	for _, x := range resident_list {
		if x.HasRequiredInfo() {
			valid_count += 1
		}
		total_count += 1
	}
	fmt.Printf("counted %d total residents, %d have the required info\n", total_count, valid_count)
	return valid_count
}

func ZeroValuesTest() {
	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}

	resident := NewResident(name, age, address)
	fmt.Printf("new resident is %v\n", resident)

	resident.Age = 0
	resident.HasRequiredInfo()
	resident.Age = 29
	resident.Address = map[string]string{"number": "404"}
	resident.HasRequiredInfo()
	resident.Address = map[string]string{"street": "Main St."}
	fmt.Printf("resident %s has the right info: %t\n", resident.Name, resident.HasRequiredInfo())
	fmt.Println("resident before delete:", resident)
	resident.Delete()
	fmt.Println("resident after delete:", resident)
	deleted_resident := resident

	name1 := "Matthew Sanabria"
	age1 := 29
	address1 := map[string]string{"street": "Main St."}
	resident1 := NewResident(name1, age1, address1)

	name2 := "Rob Pike"
	age2 := 10
	address2 := make(map[string]string)
	resident2 := NewResident(name2, age2, address2)
	residents := []*Resident{deleted_resident, resident1, resident2}

	fmt.Printf("there are %d residents w/the required info\n", CountResidents(residents))
}
