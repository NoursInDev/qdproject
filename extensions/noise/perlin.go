package noise

import (
	"math"
	"math/rand"
	//"fmt"
)

// algo ajustement transition
func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

// algo interpol linéaire = transition fluide
func lerp(t float64, a float64, b float64) float64 {
	return a + t*(b-a)
}

// algo vecteurs de gradient
func grad(hash int, x float64, y float64) float64 {
	h := hash & 7 // Gradient value 0-7
	var u float64
	if h&1 == 0 {
		u = x
	} else {
		u = -x
	}
	var v float64
	if h&2 == 0 {
		v = y
	} else {
		v = -y
	}
	return u + v
}

// algo perlin
func perlin2D(x, y float64) float64 {
	X := int(math.Floor(x)) & 255
	Y := int(math.Floor(y)) & 255

	x -= math.Floor(x)
	y -= math.Floor(y)

	u := fade(x)
	v := fade(y)

	n00 := grad(p[p[X]+Y], x, y)
	n01 := grad(p[p[X]+inc(Y)], x, y-1)
	n10 := grad(p[p[inc(X)]+Y], x-1, y)
	n11 := grad(p[p[inc(X)]+inc(Y)], x-1, y-1)

	x0 := lerp(u, n00, n10)
	x1 := lerp(u, n01, n11)

	return lerp(v, x0, x1)
}

func inc(i int) int {
	i++
	return i
}

var p []int

func initPerlin() {
	p = make([]int, 512)
	permutation := rand.Perm(256)
	for i := 0; i < 256; i++ {
		p[i] = permutation[i]
		p[256+i] = permutation[i]
	}
}

func Perlin2DMap(width, height, numBiomes int) [][]int {
	initPerlin()

	scale := 0.1
	var result [][]int

	for i := 0; i < height; i++ {
		var row []int
		for j := 0; j < width; j++ {
			value := perlin2D(float64(j)*scale, float64(i)*scale)

			biome := int(value * float64(numBiomes))

			biome++
			if biome < 0 {
				biome = 0
			} //soft patch
			row = append(row, biome)
		}
		result = append(result, row)

	}

	return result
}
