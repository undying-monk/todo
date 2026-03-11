package locking

import (
	"testing"
)

func TestAtomicCounter(t *testing.T){
	tests := []struct {// Define a struct for each test case and create a slice of them
        name string
        x int32
		numG int
        want int32
    }{
        {"10 nums", 10, 10, 20},
    }
	
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			output := RaceAtomicCounter(&tt.x, tt.numG)
			if output != tt.want{
				t.Errorf("ERROR %d", tt.x)
			}
		})
	}
}

func BenchmarkAtomicCounter(b *testing.B){
	var x int32 = 10
	for b.Loop(){
		RaceAtomicCounter(&x, 10)
	}
}