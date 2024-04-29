package quicksort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQuicksort(t *testing.T) {
    /*
    This lead to some interesting results
    I expected the use of go
    */
	arr := []int{}
    for i := 0; i < 20; i++{
        randInt := rand.Intn(11)
        arr = append(arr, randInt)
    }
    fmt.Println(arr)
    start := time.Now()
	QuicksortParallel(arr)
    duration := time.Since(start)
    fmt.Printf("Time Parallel- %v\n", duration)
    fmt.Println(arr)
    // startSingle := time.Now()
	// Quicksort(arr)
    // durationSingle := time.Since(startSingle)
    // fmt.Printf("Time Single - %v\n", durationSingle)
    //
    // fmt.Println(arr)
    // fmt.Println(arrParallel)
}
