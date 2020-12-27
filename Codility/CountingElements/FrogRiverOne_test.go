package solution

import "testing"

func TestFrogRiverOne(t *testing.T) {
	ans := FrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 5, 4})
	expected := 6

	if ans != expected {
		t.Errorf("expected %q but got %q", expected, ans)
	}
}

func BenchmarkFrogRiverOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 5, 4})
	}
}
