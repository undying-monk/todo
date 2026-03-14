package locking

import (
	"fmt"
	"testing"
)

func TestMutexCounter(t *testing.T){
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
			output := RaceMutexCounter(&tt.x, tt.numG)
			if output != tt.want{
				t.Errorf("ERROR %d", tt.x)
			}
		})
	}
}
func BenchmarkMutexCounter(b *testing.B) {
	var x int32 = 10
	var n int = 10
	for b.Loop(){
		RaceMutexCounter(&x, n)
	}
}

func BenchmarkCounterCompare(b *testing.B) {
    inputs := []int{10, 100, 1000}
	var x int32 = 10
    
    for _, n := range inputs {
        b.Run(fmt.Sprintf("MutexCounter/N=%d", n), func(b *testing.B) {
            for b.Loop() {
                RaceMutexCounter(&x, n)
            }
        })
        b.Run(fmt.Sprintf("RaceAtomicCounter/N=%d", n), func(b *testing.B) {
             for b.Loop() {
                RaceAtomicCounter(&x, n)
            }
        })
    }
}
