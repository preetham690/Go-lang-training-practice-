/*Implement using switch case
State : Karnataka , AP, Delhi ,UP
Gender: M,F
Age: >0
Height:

|State     |Gender |Height|   Age |Ticket Status|
|----------|-------|------|-------|-------------|
|Karnataka |  F    |      |       | No ticket   |
|AP        |  F    |<110cm|  <5y  | No ticket   |
|Delhi     |  F    |      |       | No Ticket   |
|UP        |  F    |<120cm|  <6y  | No ticket   |
|Karnataka |  M    |<110cm|  <5y  | No ticket   |
|AP        |  M    |<110cm|  <5y  | No ticket   |
|Delhi     |  M    |<130cm|  <7y  | No Ticket   |
|UP        |  M    |<120cm|  <6y  | No ticket   |
|Other than the above table ,It is a full ticket|
*/

package main

import "fmt"

func main() {
	busTickets("karnataka", "M", 120, 5)
}

func busTickets(state, gender string, height, age int) {
	switch {
	case (state == "karnataka" && gender == "F") || (state == "karnataka" && gender == "M" && height < 110 && age < 5):
		fmt.Println("No tickets")
	case (state == "AP" && gender == "F" && height < 110 && age < 5) || (state == "AP" && gender == "M" && height < 110 && age < 5):
		fmt.Println("No tickets")
	case (state == "delhi" && gender == "F") || (state == "delhi" && gender == "M" && height < 130 && age < 7):
		fmt.Println("No tickets")
	case (state == "UP" && gender == "F" && height < 120 && age < 6) || (state == "UP" && gender == "M" && height < 120 && age < 6):
		fmt.Println("No tickets")
	default:
		fmt.Println("Full Ticket")
	}
}
