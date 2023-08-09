package main

func main() {

	//normal for loop
	sum := 0
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			sum += i
		} else {
			sum -= i
		}

	}
	//println(sum)

	//fmt.Println(isPrime(999983))

	//temp()

	temp1()
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	i := 2
	for ; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

//--> one can implement nested for loop as well

//we can use label to break of the all the loops that are present

func temp() {
outer: // this is a label n it's user defined
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == 5 {
				break outer
			}
			println("i: ", i, "j: ", j)
		}
	}
}

//we can use goto statements as well to implement the LOOPS

func temp1() {
	i := 1
loop:
	if i%2 == 0 {
		println(i, "I is even")
	}
	i++
	if i <= 10 {
		goto loop
	}
}
