# 65. Valid Number

* https://leetcode.com/problems/valid-number/

## Fuzz Testing

Because I have not submit this problem, I cant find accepted solutions, so I collected most votes golang codes from [discuss](https://leetcode.com/problems/valid-number/discuss/?currentPage=1&orderBy=most_votes&query=&tag=golang):

* [Golang DFA, clean and understandable, 0ms beats 100%](https://leetcode.com/problems/valid-number/discuss/512431/Golang-DFA-clean-and-understandable-0ms-beats-100): Local copy [s01](s01)
* [Share my golang solution](https://leetcode.com/problems/valid-number/discuss/23915/Share-my-golang-solution): Local copy [s02](s02)
* [Share my very clear solution](https://leetcode.com/problems/valid-number/discuss/2575897/Share-my-very-clear-solution): Local copy [s03](s03)

I think their submission should be accepted before post these codes? And run [fuzzing](solution_test.go):

```
C:\joe\workspace\aitchjoe\leetcode-fuzzing\valid-number>go test -fuzz=Fuzz .
fuzz: elapsed: 0s, gathering baseline coverage: 0/1 completed
fuzz: elapsed: 0s, gathering baseline coverage: 1/1 completed, now fuzzing with 16 workers
fuzz: elapsed: 0s, execs: 14274 (61155/sec), new interesting: 31 (total: 32)
--- FAIL: FuzzIsNumber (0.24s)
    --- FAIL: FuzzIsNumber (0.00s)
        solution_test.go:44: for isNumber("4E7"), s01 - s03 = [false false true], not same

    Failing input written to testdata\fuzz\FuzzIsNumber\0a4b49cd13f63660c5c3fb035f96167c8cfd416eaf69882dc98938f988f4d951
    To re-run:
    go test -run=FuzzIsNumber/0a4b49cd13f63660c5c3fb035f96167c8cfd416eaf69882dc98938f988f4d951
FAIL
exit status 1
FAIL    github.com/aitchjoe/leetcode-fuzzing/valid-number       0.426s
```

At first I thought s03 is wrong (2 false 1 true), but after I test each code in [LeetCode](https://leetcode.com/problems/valid-number/) console, s1 and s2 failed for `"4E7"`.
