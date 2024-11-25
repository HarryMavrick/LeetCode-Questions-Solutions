package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	for i := 0; i <= len(strs[0])-1; i++ {
		char := strs[0][i]

		for _, values := range strs[1:] {
			if i >= len(values) || values[i] != char {
				return strs[0][:i]
			}
		}
	}
	return strs[0]

}
func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println("The longest common prifix is- ", result)
}
