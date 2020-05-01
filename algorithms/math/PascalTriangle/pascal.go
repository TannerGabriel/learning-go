package PascalTriangle

func pascalTriangle(height int) [][]int {
	pascalTriangle := make([][]int, height)

	// Make pacal triangle row
	for i := 0; i < height; i++ {
		pascalTriangle[i] = make([]int, i+1)

		pascalTriangle[i][0] = 1

		// Fill row with values
		for j := 1; j < i; j++ {
			pascalTriangle[i][j] = pascalTriangle[i-1][j] + pascalTriangle[i-1][j-1]
		}

		pascalTriangle[i][i] = 1
	}

	return pascalTriangle
}
