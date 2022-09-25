package solution

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/aitchjoe/leetcode-fuzzing/median-of-two-sorted-arrays/s01"
	"github.com/aitchjoe/leetcode-fuzzing/median-of-two-sorted-arrays/s02"
	"github.com/aitchjoe/leetcode-fuzzing/median-of-two-sorted-arrays/s03"
)

type TestCase struct {
	Input1 []int
	Input2 []int
	Want  float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	testcases := []TestCase{
		{
			Input1: []int{1, 3},
			Input2: []int{2},
			Want:  2.00000,
		},
		{
			Input1: []int{1, 2},
			Input2: []int{3, 4},
			Want:  2.50000,
		},
	}
	for _, tc := range testcases {
		out := findMedianSortedArrays(tc.Input1, tc.Input2)
		if out != tc.Want {
			t.Fatalf(`findMedianSortedArrays(%v, %v) = %f, want %f`, tc.Input1, tc.Input2, out, tc.Want)
		}
	}
}

func FuzzFindMedianSortedArrays(f *testing.F) {
	rand.Seed(time.Now().UnixNano())
	f.Add(10, 20)
	f.Add(0, 999)
	f.Fuzz(func(t *testing.T, m, n int) {
		if m < 0 {
			m = -m
		}
		m %= 1001
		if n < 0 {
			n = -n
		}
		n %= 1001
		if m + n == 0 {
			return
		}

		nums1 := make([]int, m, m)
		for i := 0; i < m; i++ {
			nums1[i] = int(rand.Uint32() % 1000001)
		}
		sort.Ints(nums1)
		nums2 := make([]int, n, n)
		for i := 0; i < n; i++ {
			nums2[i] = int(rand.Uint32() % 1000001)
		}
		sort.Ints(nums2)

		my := findMedianSortedArrays(nums1, nums2)
		ss := make([]float64, 3, 3)
		ss[0] = s01.FindMedianSortedArrays(nums1, nums2)
		ss[1] = s02.FindMedianSortedArrays(nums1, nums2)
		ss[2] = s03.FindMedianSortedArrays(nums1, nums2)
		same := true
		for _, s := range ss {
			if s != my {
				same = false
				break
			}
		}

		if !same {
			t.Errorf(`for findMedianSortedArrays(%v, %v), my = %f, s01 - s0%d = %v, not same`, nums1, nums2, my, len(ss), ss)
		}
	})
}
