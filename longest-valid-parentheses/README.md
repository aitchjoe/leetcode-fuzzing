# 32. Longest Valid Parentheses

* https://leetcode.com/problems/longest-valid-parentheses/

## Fuzz Testing

All passed for 17 minutes fuzzing:
```
C:\joe\workspace\aitchjoe\leetcode-fuzzing\longest-valid-parentheses>go test -fuzz=Fuzz .
fuzz: elapsed: 0s, gathering baseline coverage: 0/96 completed
fuzz: elapsed: 0s, gathering baseline coverage: 96/96 completed, now fuzzing with 16 workers
fuzz: elapsed: 3s, execs: 286002 (95302/sec), new interesting: 8 (total: 104)
fuzz: elapsed: 6s, execs: 574796 (96090/sec), new interesting: 9 (total: 105)
......
fuzz: elapsed: 17m21s, execs: 92166111 (79480/sec), new interesting: 14 (total: 96)
fuzz: elapsed: 17m24s, execs: 92389884 (74520/sec), new interesting: 14 (total: 96)
<Ctrl+C>
PASS
ok      github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses  1044.399s
```
