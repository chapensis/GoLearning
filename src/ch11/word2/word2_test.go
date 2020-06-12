package word

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayaK", true},
		{"detartrated", true},
		{"a man ", false},
		{"asdasdsddw", false},
		{"desserts", false},
	}
	for _, test := range tests{
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func randomPalindrome(rng *rand.Rand) string{
	n := rng.Intn(25) // 随机字符的最大长度是24
	runes := make([]rune, n)
	for i := 0; i < (n + 1) / 2; i++ {
		r := rune(rng.Intn(0x1000)) // 随机字符最大是'\u0999'
		runes[i] = r
		runes[n - i - 1] = r
	}
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	// 初始化一个伪随机生成器
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed:%d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}