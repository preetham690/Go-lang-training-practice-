package main

func main() {
	whichDay(4)
	isValid(0)
	isDivisible(28)
	fallNegative(28)
}

// switch statement
func whichDay(day int) {
	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thr")
	case 6:
		println("Fri")
	case 7:
		println("Sat")
	default:
		println("Invalid day")
	}
}

//empty switch statement

func isValid(num int) {
	switch {
	case num >= 0 && num < 50:
		println("it's between 0-50")
	case num >= 50 && num < 100:
		println("Its between 50-100")
	case num >= 100 && num < 200:
		println("it's between 100-200")
	default:
		println("above 200")
	}
}

//also cases can have multiple conditions

// FALL THROUGH -> Keyword

// removing break is equal to giving fall through
// this can leads to fall negative as well, so we need to be careful while using fallthrough
func isDivisible(num int) {
	switch {
	case num%8 == 0:
		println("It is divisible by 8")
		fallthrough
	case num%4 == 0:
		println("It is divisible by 4")
		fallthrough
	case num%2 == 0:
		println("It is divisible by 2")
	default:
		println("Not divisible by 2,4 8")
	}
}

//below example demostrates the fall negative

func fallNegative(num int) {
	switch {
	case num%2 == 0:
		println("It is divisible by 2")
		fallthrough
	case num%4 == 0:
		println("It is divisible by 4")
		fallthrough
	case num%8 == 0:
		println("It is divisible by 8")
		fallthrough
	default:
		println("Not divisible by 2,4 8")
	}
}
