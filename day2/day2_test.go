package day2

import(
	"testing"
	"github.com/stretchr/testify/require"
)

func TestQ2(t *testing.T) {
	ans1 := RunStar1("test1.txt")
	require.Equal(t, 15, ans1)

	ans2 := RunStar2("test1.txt") 
	require.Equal(t, 12, ans2)
}