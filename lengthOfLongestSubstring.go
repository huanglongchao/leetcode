package main

func lengthOfLongestSubstring(s string) int{

	var maxL = 0
	var i,j = 0, 0
	set := make(map[string]string)

	for i < len(s) && j < len(s) {
		_, ok := set[s[j:j+1]]
		if !ok {
			set[s[j:j+1]] = ""
			maxL = max(maxL, j-i+1)
			j++
		} else {
			delete(set, s[i:i+1])
			i++
		}
	}
	return maxL
}

func max(a,b int) int{
	if a > b{
		return a
	}
	return b
}


