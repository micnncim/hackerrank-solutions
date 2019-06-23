# [Minimum Swaps 2](https://www.hackerrank.com/challenges/minimum-swaps-2/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays)

## メモ

- パターンは、好ましいものの順に、
  - インデックスと値が等しいため、swap する必要が無い
  - swap すると同時にインデックスと値が等しくなる2つの数字がある
    - 2数字につき1 swap
    - 最後の2数字は必ずこのパターン
  - swap すると一方はインデックスと値が等しくなるが、他方は等しくならない
    - 1数字につき1 swap
  