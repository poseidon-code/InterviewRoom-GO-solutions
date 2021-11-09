// PROBLEM: Find median in a stream
// https://practice.geeksforgeeks.org/problems/find-median-in-a-stream-1587115620/1

package main

import (
	"fmt"
	"math"
)

var median float32
var maxq, minq []int

// SOLUTION
func balance_heaps() {
    if math.Abs(float64(len(maxq) - len(minq))) > 1 {
        if len(maxq) > len(minq) {
            minq = append(minq, maxq[0])
            maxq = maxq[1:]
        } else {
            maxq = append(maxq, minq[0])
            minq = minq[1:]
        }
    }
}

func insert_heap(x int) {
    if len(maxq)==0 {
        maxq = append(maxq, x)
    } else if x > maxq[0] {
        minq = append(minq, x)
    } else if len(minq)==0 {
        minq = append(minq, maxq[0])
        maxq = maxq[1:]
        maxq = append(maxq, x)
    } else {
        maxq = append(maxq, x)
    }

    balance_heaps()
}

func get_median() float32 {
    if ((len(maxq) + len(minq)) % 2) == 0 {
        median = float32(maxq[0] + minq[0])
        median /= 2
    } else {
        if len(minq)==0 {
            median = float32(maxq[0])
        } else if len(minq) < len(maxq) {
            median = float32(maxq[0])
        } else {
            median = float32(minq[0])
        }
    }

    return median
}


func main() {
    // INPUT :
    x := []int{5,10,15}
    

    // OUTPUT :
    for _, i := range x {
        insert_heap(i)
        fmt.Println(get_median())
    }
}