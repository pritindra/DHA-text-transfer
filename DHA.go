package main

import (
	"math"
	"math/rand"
	"time"
)

var generator float64 = 2
var prime float64 = 600
var key_length int = 23

// func power(a ,b , p) int{
//   if b == 1 {
//     return a
//   } else {
//     return math.Pow(a,b)% p
//   }
//
// }

func generate_private_key(length int) int64 {
	rand.Seed(time.Now().UnixNano())
	// _rand := 0
	_bytes := length // 8 + 8
	private_key := rand.Intn(_bytes)
	return int64(private_key)
}

func generate_public_key(private_key int) int64 {
	public_key := int64(math.Pow(generator, float64(private_key)))
	return public_key
}

func generate_secret_key(private_key int64, public_key int64) int64 {
	PowerFunc := math.Pow(float64(public_key), float64(private_key))
	secret_key := math.Mod(PowerFunc, prime)
	return int64(secret_key)
}

func main() {
	generate_private_key(10)
}
