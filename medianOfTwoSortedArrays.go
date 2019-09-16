package main

import (
	"math"
)

/**
	Example 2:

	nums1 = [1, 2]
	nums2 = [3, 4]

	The median is (2 + 3)/2 = 2.5
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
		m = len(nums1)
		n = len(nums2)
	}
	var imin,imax = 0,m
	var i,j = 0,0
	var maxLeft,minRight = 0.0,0.0
	if m == 0 && n == 0{
		return 0
	}else if m ==0 && n == 1{
		return float64(nums2[0])
	}else{
		for imin <= imax {
			i = (imin + imax)/2
			j = (m+n+1)/2 - i
			if i < m && nums2[j-1] > nums1[i]{
				imin = i + 1
			}else if i > 0 && nums1[i-1] > nums2[j]{
				imax = i - 1
			}else{
				if i == 0{
					maxLeft = float64(nums2[j-1])
				}else if j ==0{
					maxLeft = float64(nums1[i-1])
				}else{
					maxLeft = math.Max(float64(nums1[i-1]),float64(nums2[j-1]))
				}

				if i == m{
					minRight = float64(nums2[j])
				}else if j ==n {
					minRight = float64(nums1[i])
				}else{
					minRight = math.Min(float64(nums2[j]),float64(nums1[i]))
				}
				break
			}
		}
		if (m+n)%2 == 0{
			return (maxLeft + minRight) /2.0
		}else{
			return maxLeft
		}
	}
}