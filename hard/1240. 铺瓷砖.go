package main

// https://leetcode.cn/problems/tiling-a-rectangle-with-the-fewest-squares/

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 *                      widthLong
 *              |-----------------------|
 *            - ######################### -
 * heightLong | ######################### | heightShort
 *            | ######################### -
 *            | ##########
 *            | ##########
 *            - ##########
 *              |--------|
 *                      widthShort
 */

func tilingRectangleSolve(heightLong, heightShort, widthLong, widthShort int) int {
	//截止条件
	if heightLong == 0 || widthLong == 0 {
		return 0
	}
	//正方形
	if heightLong == heightShort && widthLong == widthShort && heightLong == widthLong {
		return 1
	}
	var ans = 10000
	var temp int
	var maxDim = minInt(heightShort, widthShort)
	if heightLong == heightShort {
		//当是个矩形的时候
		for i := maxDim; i > 0; i-- {
			//还是矩形
			if heightLong-i == 0 {
				temp = 1 + tilingRectangleSolve(heightLong, heightLong, widthLong-i, widthLong-i)
			} else if widthLong-i == 0 {
				temp = 1 + tilingRectangleSolve(heightLong-i, heightLong-i, widthLong, widthLong)
			} else {
				temp = 1 + tilingRectangleSolve(heightLong, heightLong-i, widthLong, widthLong-i)
			}
			ans = minInt(ans, temp)
		}
	} else {
		//当不是个矩形的时候
		if maxDim == heightShort {
			if widthLong-maxDim > widthShort {
				temp = 1 + tilingRectangleSolve(heightLong, heightShort, widthLong-maxDim, widthShort)
			}
			if widthLong-maxDim == widthShort {
				temp = 1 + tilingRectangleSolve(heightLong, heightLong, widthShort, widthShort)
			}
			if widthLong-maxDim < widthShort {
				temp = 1 + tilingRectangleSolve(heightLong, heightLong-maxDim, widthShort, widthLong-maxDim)
			}
		} else {
			if heightLong-maxDim > heightShort {
				temp = 1 + tilingRectangleSolve(heightLong-maxDim, heightShort, widthLong, widthShort)
			}
			if heightLong-maxDim == heightShort {
				temp = 1 + tilingRectangleSolve(heightShort, heightShort, widthLong, widthLong)
			}
			if heightLong-maxDim < heightShort {
				temp = 1 + tilingRectangleSolve(heightShort, heightLong-maxDim, widthLong, widthLong-maxDim)
			}
		}
		ans = minInt(ans, temp)
	}
	return ans
}

func tilingRectangle(n int, m int) int {
	return tilingRectangleSolve(n, n, m, m)
}
