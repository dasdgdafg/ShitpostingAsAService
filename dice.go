package main

import (
	"crypto/rand"
	"math"
	"math/big"
)

func rollDice(sidesPerDie int64, numberOfDice int64) string {
	if numberOfDice > 1000000 {
		return rollApprox(sidesPerDie, numberOfDice).String()
	} else {
		return rollExact(sidesPerDie, numberOfDice).String()
	}
}

func rollExact(sidesPerDie int64, numberOfDice int64) *big.Int {
	e := int64(8) // from experimentation, 8 seems to work well
	sides := big.NewInt(sidesPerDie)
	// calculate numberOfDice/e d sidesPerDie^e, instead of numberOfDice d sidesPerDie, ie 2d36 instead of 4d6
	// becuase 1 large random number seems to be faster than 2 small random numbers
	nn := big.NewInt(0).Exp(sides, big.NewInt(e), nil)
	mm := numberOfDice / e
	sum := big.NewInt(0)
	temp := big.NewInt(0) // declared out here to avoid allocating a new one each iteration of the loop
	for i := int64(0); i < mm; i++ {
		roll, err := rand.Int(rand.Reader, nn)
		if err != nil {
			panic(err)
		}
		// convert the number back into the individual rolls
		// ie, 1d36 = 36 -> 2d6 = 6,6
		for j := int64(1); j < e; j++ {
			roll.DivMod(roll, sides, temp)
			sum.Add(sum, temp)
		}
		sum.Add(sum, roll)
	}
	// if numberOfDice wasn't divisible by e, compute any extra rolls 1 at a time
	for i := int64(0); i < numberOfDice%e; i++ {
		roll, err := rand.Int(rand.Reader, sides)
		if err != nil {
			panic(err)
		}
		sum.Add(sum, roll)
	}
	// the random numbers are in [0,n-1] instead of [1,n], so add 1 to the sum for each roll that happened
	sum.Add(sum, big.NewInt(numberOfDice))
	return sum
}

func randFloat() float64 {
	// returns a random float in the range (0,1)
	// crypto/rand seems to only do ints, so divide by MaxInt64
	n_int, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64-1))
	if err != nil {
		panic(err)
	}
	n_int.Add(n_int, big.NewInt(1)) // don't let the number be 0
	n := new(big.Float).SetInt(n_int)
	n.Quo(n, new(big.Float).SetInt64(math.MaxInt64))
	n_float, _ := n.Float64() // don't care about the accurracy
	return n_float
}

func rollApprox(sidesPerDie int64, numberOfDice int64) *big.Int {
	// approximate using the Box-Muller method
	u := randFloat()
	v := randFloat()
	x := big.NewFloat(math.Sqrt(-2.0*math.Log(u)) * math.Cos(2*math.Pi*v))

	// find the mean and standard deviation
	num := new(big.Float).SetInt64(numberOfDice)
	// mean = (1+sidesPerDie)/2 * numberOfDice
	mean := new(big.Float).SetInt64(1)
	mean.Add(mean, new(big.Float).SetInt64(sidesPerDie))
	mean.Mul(mean, num)
	mean.Quo(mean, new(big.Float).SetInt64(2))
	// variance = ((max - min + 1)^2 - 1) / 12 * numberOfDice
	// min is always 1 here
	variance := new(big.Float).SetInt64(sidesPerDie)
	variance.Mul(variance, variance)
	variance.Sub(variance, new(big.Float).SetInt64(1))
	variance.Quo(variance, new(big.Float).SetInt64(12))
	variance.Mul(variance, num)
	stddev := new(big.Float).Sqrt(variance)

	// convert to the desired distribution
	// x * stddev + mean
	x = stddev.Mul(x, stddev)
	x.Add(x, mean)
	// convert to an int
	x_int := new(big.Int)
	x.Int(x_int)
	return x_int
}
