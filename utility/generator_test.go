package utility

import "testing"

func TestGenerateRandom(t *testing.T) {
	tests := []struct{
		inputSize int
		expectedSize int
	}{
		{2 , len(generateList(2))},
	}

	for _,test := range tests {
		if test.expectedSize != test.inputSize {
			t.Errorf("Expected sizes are not the same || input = %d | expected = %d", test.inputSize, test.expectedSize)
		}
	}
}

func BenchmarkGenerateRandom(b *testing.B) {
	for i:=0; i < b.N; i++ {
		generateList(100)
	}
}