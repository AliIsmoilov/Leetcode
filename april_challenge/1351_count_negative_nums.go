package aprilchallenge

// Example 1:

// Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
// Output: 8
// Explanation: There are 8 negatives number in the matrix.
// Example 2:

// Input: grid = [[3,2],[1,0]]
// Output: 0

func CountNegatives(grid [][]int) int {
    count := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            // fmt.Println(grid[i][j])
            if grid[i][j] < 0 {
                count++
            }
        }
    }

    return count
}