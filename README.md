# LeetCode Fuzz Testing

## Background

After my [submission](https://leetcode.com/submissions/detail/807576266/) of [Longest Subarray of 1's After Deleting One Element](https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/) is accepted, I found a bug in my code (check [this](https://github.com/aitchjoe/leetcode-fuzzing/commit/4b6bf70966f7c53adf1a7e089b59772de73c95d3) commit), but just by think, not by test (74 / 74 test cases passed).

At the same time, I am studying [Go Fuzzing](https://go.dev/security/fuzz/) (fuzz testing). In most complex situations, the hardest part of fuzzing is how to verify the result for any input, but I found if LeetCode use fuzzing, the problem change to very easy, because we dont need to verify the result is correct or not, we just run all accepted submissions of the same problem and compare all results, if not same, then we find the new test case.

I have collected some source codes from Accepted Solutions Runtime Distribution and Accepted Solutions Memory Distribution (I have not found a easy way to do this), run such as [FuzzLongestSubarray()](longest-subarray-of-1s-after-deleting-one-element/solution_test.go), then I found the same error in other's [accepted submission](longest-subarray-of-1s-after-deleting-one-element#fuzz-testing). Although it is same error, but the important difference IS that we dont need to go deep to find bugs or create more edge test cases, the fuzzing show these to us.

I will try to fuzzing other problems. TODO

## Test

Under each problem directory, run:

* `go test -v .`
* `go test -fuzz=Fuzz .`

## Problems

* [1493. Longest Subarray of 1's After Deleting One Element](longest-subarray-of-1s-after-deleting-one-element)
