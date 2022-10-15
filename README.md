# LeetCode Fuzz Testing

## Background

After my [submission](https://leetcode.com/submissions/detail/807576266/) of [Longest Subarray of 1's After Deleting One Element](https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/) is accepted, I found a bug in my code (check [this](https://github.com/aitchjoe/leetcode-fuzzing/commit/4b6bf70966f7c53adf1a7e089b59772de73c95d3) commit), but just by think, not by test (74 / 74 test cases passed).

At the same time, I am studying [Go Fuzzing](https://go.dev/security/fuzz/) (fuzz testing). In most complex situations, the hardest part of fuzzing is how to verify the result for any input, but I found if LeetCode use fuzzing, the problem change to very easy, because we dont need to verify the result is correct or not, we just run all accepted submissions of the same problem and compare all results, if not same, then we find the new test case.

I have collected some source codes from Accepted Solutions Runtime Distribution and Accepted Solutions Memory Distribution (I have not found a easy way to do this), run such as [FuzzLongestSubarray()](longest-subarray-of-1s-after-deleting-one-element/solution_test.go), then I found the same error in other's [accepted submission](longest-subarray-of-1s-after-deleting-one-element#fuzz-testing). Although it is same error, but the important difference IS that we dont need to go deep to find bugs or create more edge test cases, the fuzzing show these to us. I will do more fuzzing on [Problems](#problems).

## Test

Under each problem directory, run:

* `go test -v .`
* `go test -fuzz=Fuzz -fuzztime=55m .`

Check the results in [GitHub Actions](https://github.com/aitchjoe/leetcode-fuzzing/actions).

## Problems

* [4. Median of Two Sorted Arrays](median-of-two-sorted-arrays)
* [32. Longest Valid Parentheses](longest-valid-parentheses)
* [65. Valid Number](valid-number)
* [1493. Longest Subarray of 1's After Deleting One Element](longest-subarray-of-1s-after-deleting-one-element)

After fuzzing problem 1493 and 4, no more interesting bugs found, so I look for hard and lowest acceptance rate [problems](https://leetcode.com/problemset/all/?difficulty=HARD&page=1&sorting=W3sic29ydE9yZGVyIjoiQVNDRU5ESU5HIiwib3JkZXJCeSI6IkFDX1JBVEUifV0%3D), then I have some progress on problem [65](valid-number).
