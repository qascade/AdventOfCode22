package day1

import(
	"testing"
	"github.com/stretchr/testify/require"
)

func TestQ1(t *testing.T) {
	ans1 := RunStar1("test1.txt")
	require.Equal(t, 6, ans1)

	ans2 := RunStar2("test1.txt") 
	require.Equal(t, 10, ans2)
}