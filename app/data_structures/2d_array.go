package datastructures

func New2DArray[T any](rows, cols int, initVal T) [][]T {
	arr := make([][]T, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]T, cols)
		for j := 0; j < cols; j++ {
			arr[i][j] = initVal
		}
	}
	return arr
}

func PrettyPrint[T any](arr [][]T) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			print(arr[i][j], " ")
		}
		println()
	}
}
