package main

import "fmt"

type Accounting struct {
}

func (Accounting) WorkWithExcell() string {
	return "Accounting on work now, please don't disturb"
}

type StaffAdmin struct {
}

func (StaffAdmin) WriteDailyReport() string {
	return "admin write daily report"
}

func (StaffAdmin) WriteScheduleForMeet() string {
	return "admin write schedule for meet"
}

func main() {
	acc := Accounting{}
	fmt.Printf("acc.WorkWithExcell(): %v\n", acc.WorkWithExcell())

	adm := StaffAdmin{}
	fmt.Printf("adm.WriteDailyReport(): %v\n", adm.WriteDailyReport())
	fmt.Printf("adm.WriteScheduleForMeet(): %v\n", adm.WriteScheduleForMeet())
}
