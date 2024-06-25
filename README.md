| 티켓 이름 | 가격 |
|-------|----|
| A     | 1  |
| B     | 2  |
| c     | 3  |
| d     | 4  |

<상점의 티켓 진열 예측>

| 상점 순서   | 진열 티켓      |
|---------|------------|
| 최초 상점   | B, C, B, C |
| 두 번째 상점 | A, A, A, B |
| 세 번째 상점 | D, D, C, D |

| tickets                      | roll | shop                                                                 | money | result |
|------------------------------|------|----------------------------------------------------------------------|-------|--------|
| ["A 1", "B 2", "C 5", "D 3"] | 10   | [["B", "C", "B", "C"], ["A", "A", "A", "B"], ["D", "D", "C", "D"]]   | 30    | 2      
| ["A 1", "B 2", "C 5", "D 3"] | 10   | [["B", "C", "B", "C"], ["A", "A", "A", "B"], ["D", "D", "C", "D"]]   | 100   | 4      
| ["A 2", "B 1"]               | 1    | [["A", "A", "A"], ["A", "B", "B"], ["A", "B", "B"], ["A", "B", "B"]] | 9     | 2      
| ["A 2", "B 1"]               | 5    | [["A", "A", "A"], ["A", "B", "B"], ["A", "B", "B"], ["A", "B", "B"]] | 19    | 2      