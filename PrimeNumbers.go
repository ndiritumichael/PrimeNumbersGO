package main
//calculating prime numbers using sieve of eratosthenes

import (
	"fmt"
	"time"
)

func PrimeNumbers(x int)[]int{
	//function takes in an upper limit number to calculate from zeo and returns a slice of containing prime numbers in the range

	numbers:= make([]bool,x+1)   //create an slice of booleans equal to the size total numbers plus 1
	for i:=2;i<x+1;i++{

		numbers[i] = true //initialize every index as true

	}
	for prime:=2;prime*prime<=x;prime ++{
		if numbers[prime]== true{
			for i :=prime *2;i<=x;i+=prime{
				numbers[i] = false

			}
		}
	}
	var Result []int
	for i:=2;i<=x;i++{
		if numbers[i]==true{
			Result = append(Result,i)
		}
	}
	return Result //return the slice of prime numbers

}
func main(){
	start := time.Now()
	x :=PrimeNumbers(1000000000)
	fmt.Println(x)
	elapsed := time.Since(start)
	fmt.Printf("1M has %d primes calculated in %s milliseconds",len(x),elapsed)


}
