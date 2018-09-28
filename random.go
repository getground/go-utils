package goutils


import (
    "math/rand"
)


// see https://flaviocopes.com/go-random/

// Returns an int >= min, < max
func RandomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func RandomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(RandomInt(97, 122)) // lowercase only
    }
    return string(bytes)
}
