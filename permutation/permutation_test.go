package permutation

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}

func pnk(list []int, k int) int {
	return factorial(len(list)) / factorial(len(list)-k)
}

func TestGenerateRepetitionsFalse(t *testing.T) {
	list := randomList()

	// Check if the correct number of permutations is generated
	for k := 0; k <= len(list); k++ {
		assert.Equal(t, pnk(list, k), len(Generate(list, &k, false)))
	}

	// Check if permutations have no repetitions
	for k := 0; k <= len(list); k++ {
		actualPerms := Generate(list, &k, false)

		for _, perm := range actualPerms {
			seen := make(map[int]bool)
			for _, element := range perm {
				_, exists := seen[element]
				assert.False(t, exists)
				seen[element] = true
			}
		}
	}
}

func TestGenerateRepetitionsTrue(t *testing.T) {
	list := randomList()
	for k := 0; k <= len(list); k++ {
		assert.Equal(t, len(Generate(list, &k, true)), intPow(len(list), k))
	}
}

func randomList() []int {
	length := 5
	list := make([]int, length)

	for i := 0; i < length; i++ {
		list[i] = rand.Intn(100)
	}

	return list
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
