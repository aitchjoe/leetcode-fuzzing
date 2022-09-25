package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums := []int{}
    i := 0
    j := 0
    for i < len(nums1) || j < len(nums2) {
        if j < len(nums2) && ( i == len(nums1) || nums2[j] < nums1[i]) {
            nums = append(nums, nums2[j])
            j++
        } else if i < len(nums1) {
            nums = append(nums, nums1[i])
            i++
        }
    }
    l := len(nums)
    if l % 2 == 1 {
        return float64(nums[(l + 1) / 2 - 1])
    }
    a := nums[l / 2 - 1]
    b := nums[l / 2]
    return float64((a + b)) / 2
}
