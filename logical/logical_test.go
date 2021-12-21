package logical

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAndNot(t *testing.T) {
	r := andNot(1, 0)
	fmt.Println("1 &^ 0: ", r)
	assert.Equal(t, 1, r)

	r = andNotNot(1, 0)
	fmt.Println("1 &^ (^0): ", r)
	assert.Equal(t, 0, r)

	fmt.Println("1 &^ (^1): ", andNotNot(1, 1))
	fmt.Println("0 &^ (^1): ", andNotNot(0, 1))
	fmt.Println("0 &^ (^0): ", andNotNot(0, 0))
}
