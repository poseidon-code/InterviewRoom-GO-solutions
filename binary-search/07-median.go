// PROBLEM: Median of two sorted arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/
// https://www.interviewbit.com/problems/median-of-array/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func median(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    l, r := 0, m
    if m > n {return median(nums2, nums1)}

    for l <= r {
        i := (l+r)/2
        j := (m+n+1)/2 - i

        if i!=0 && nums1[i-1] > nums2[j] {
            r = i-1
        } else if i<m && nums2[j-1] > nums1[i] {
            l = i+1
        } else {
            lmax := func()int{
                if i==0 {
                    return nums2[j-1]
                } else {
                    if j==0 {
                        return nums1[i-1]
                    } else {
                       return int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
                    }
                }
            }()

            if ((m + n)%2)!=0 {return float64(lmax)}

            rmin := func()int{
                if i==m {
                    return nums2[j]
                } else {
                    if j==n {
                        return nums1[i]
                    } else {
                       return int(math.Min(float64(nums1[i]), float64(nums2[j])))
                    }
                }
            }()

            return 0.5 * float64((lmax + rmin))
        }
    }

    return 0.0
}


func main() {
    // INPUT :
    nums1 := []int{1,2}
    nums2 := []int{3,4}

    // OUTPUT :
    fmt.Println(median(nums1, nums2))
}