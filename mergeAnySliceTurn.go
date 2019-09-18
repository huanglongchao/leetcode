package main

import "fmt"

func merge(ss ...[]int)[]int{
	var res []int
	maxL := 0
	for _,s := range ss{
		if len(s) > maxL{
			maxL = len(s)
		}
	}
	for i:=0; i< maxL; i++  {
		for _,s := range ss {
			if i < len(s){
				res = append(res,s[i])
			}
		}
	}
	return res
}

func main(){
	fmt.Println(merge([]int{1,2,3},[]int{4,5,6},[]int{7,8}))
}
