# [Arrays: Left Rotation](https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays&h_r=next-challenge&h_v=zen)

## メモ

- 配列を左に N 個ズラす場合は、(`元のインデックス` + `ズラす数`) % `配列の長さ` のインデックスの値になる
  - e.g.) 2個ズラす場合: [1, 2, 3, 4, 5] -> [3, 4, 5, 1, 2]
    - 3 = a[(0+2)%5=2]
    - 4 = a[(1+2)%5=3]
    - 5 = a[(1+3)%5=4]
    - 1 = a[(1+4)%5=0]
    - 2 = a[(1+5)%5=1]
