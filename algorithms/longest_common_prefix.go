package algorithms

/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
	Input: strs = ["flower","flow","flight","flamming"]
	Output: "fl"

Example 2:
	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.

Constraints:
	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] consists of only lowercase English letters if it is non-empty.
*/

func LongestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(prefix) >= len(strs[i]) {
			prefix = prefix[:len(strs[i])]
		}
		word := strs[i]
		for i := len(prefix) - 1; i >= 0; i-- {
			if prefix[i] != word[i] {
				prefix = prefix[:i]
			}
		}
	}
	return prefix
}
