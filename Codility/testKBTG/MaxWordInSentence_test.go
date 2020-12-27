package testKBTG

import "testing"

func TestFindMaxWord(t *testing.T) {

	stringTests := []struct {
		text string
		max  int
	}{
		{"This is testCoder, Give me a job?", 4},
		{"This is testCoder, Give me a job          ? ", 4},
		{"This is testCoder, Give me a job?    ", 4},
		{"This is testCoder , Give me a job?", 4},
		{" This is testCoder  ,   Give me a job?  ", 4},
		{"    This is testCoder    , Give me a job   ?     ", 4},
		{"  Thisis testCode r    ,  Give   me a job  ?    ", 4},
		{"  This  ist?estCo   der   , G i v !em e!     a asdad    adsdasdas  job ? ! !   ,,,!?,  ,! ,? ,, ?? !! ", 4},
		{"T!!!h!s is tes????tCode???r, , , Give me a job ? sad x.x, x  ! , x.x.x x w.22 dassa   2   D ???", 6},
		{"T!!!h!s is tes????tCode???r, , , Give me a job ? sad x.x, x  ! , x.x.x x w.22 dassa   2   D ???", 6},
	}

	for i, tt := range stringTests {
		t.Run("test1", func(t *testing.T) {
			got := FindMaxWord(tt.text)
			if got != tt.max {
				t.Errorf("test: %d got %d want %d", i, got, tt.max)
			}
		})
	}

}

func BenchmarkFindMaxWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxWord("This is testCoder, Give me a job?")
	}
}
