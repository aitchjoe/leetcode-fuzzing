# 1493. Longest Subarray of 1's After Deleting One Element

* https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

## Fuzz Testing

([s02](s02)) Accepted Solutions Memory Distribution - [sample 7200 KB submission](https://leetcode.com/submissions/api/detail/1586/golang/memory/7200/) failed:

```
C:\joe\workspace\aitchjoe\leetcode-fuzzing\longest-subarray-of-1s-after-deleting-one-element>go test -fuzz=Fuzz .
fuzz: elapsed: 0s, gathering baseline coverage: 0/62 completed
fuzz: elapsed: 0s, gathering baseline coverage: 14/62 completed
--- FAIL: FuzzLongestSubarray (0.07s)
    --- FAIL: FuzzLongestSubarray (0.00s)
        solution_test.go:83: for longestSubarray([1 0]), my = 1, s01 - s05 = [1 0 1 1 1], not same

    Failing input written to testdata\fuzz\FuzzLongestSubarray\319194200cbebc1b28cd91ec5abfe2914d02e370e20c0c2495f6a9c559ba26a3
    To re-run:
    go test -run=FuzzLongestSubarray/319194200cbebc1b28cd91ec5abfe2914d02e370e20c0c2495f6a9c559ba26a3
FAIL
exit status 1
FAIL    github.com/aitchjoe/leetcode-fuzzing/longest-subarray-of-1s-after-deleting-one-element  0.266s
```
