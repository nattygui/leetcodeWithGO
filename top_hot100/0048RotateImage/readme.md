# 48. 旋转图像

给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

```
给定 matrix = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
```

示例 2:

```
给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
], 

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
```

## 解法

图像的大小为 **NxN**

将图像顺时针旋转90度，通过观察可以看出原矩阵martix1中的某一行 **i** 和某一列 **j** 和旋转后的矩阵martix2的关系为

- martrix1[i][j] = martrix2[j][n-i]

同理有
- martrix2[i][j] = martrix1[j][n-i]

```go
func rotate(matrix [][]int) {
	size := len(matrix)
	subsize := size - 1
	for i := 0; i < size/2; i++ {
		for j := i; j < subsize-i; j++ {
			matrix[i][j], matrix[j][subsize-i] = matrix[j][subsize-i], matrix[i][j]
			matrix[i][j], matrix[subsize-i][subsize-j] = matrix[subsize-i][subsize-j], matrix[i][j]
			matrix[i][j], matrix[subsize-j][i] = matrix[subsize-j][i], matrix[i][j]
		}
	}
}
```