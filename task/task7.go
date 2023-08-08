/*State : Karnataka , AP, Delhi ,UP
Gender: M,F
Age: >0
Height:

            Gender  Height   Age   Ticket Status
Karnataka    F                     No ticket
AP           F      <110cm   <5y   No ticket
Delhi        F                     No Ticket
UP           F      <120cm   <6y   No ticket

Karnataka    M      <110cm    <5y  No ticket
AP           M      <110cm    <5y  No ticket
Delhi        M      <130cm    <7y  No Ticket
UP           M      <120cm    <6y  No ticket

Other than the above table , It is full ticket*/

package main

import "fmt"

func main() {
	generateTicket("Karnataka", "M", 22, 182)
	generateTicket("AP", "F", 3, 70)
	generateTicket("Delhi", "M", 6, 100)
	generateTicket("UP", "F", 5, 130)
	generateTicket("Kerla", "M", 3, 90)
}

func generateTicket(state, gender string, age, height int) {

	if state == "Karnataka" {
		if (gender == "F") || (gender == "M" && height < 110 && age < 5) {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}
	} else if state == "AP" {
		if (gender == "F" && height < 110 && age < 5) || (gender == "M" && height < 110 && age < 5) {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}
	} else if state == "Delhi" {
		if (gender == "F") || (gender == "M" && height < 130 && age < 7) {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}
	} else if state == "UP" {
		if (gender == "F" && height < 120 && age < 6) || (gender == "M" && height < 120 && age < 6) {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}
	} else { // this condition is picked when the state is diff
		fmt.Println("Full Ticket")
	}

}
