// 2187. 完成旅途的最少时间
// URL：https://leetcode-cn.com/problems/minimum-time-to-complete-trips/
// 难度：中等
// 关键词：数学、二分搜索
// 执行用时：192 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：8.5 MB, 在所有 Go 提交中击败了 100.00% 的用户

package main

import "math"

func minimumTime(time []int, totalTrips int) int64 {
	var unitTrips, lowerBound, upperBound float64

	trips := func(t int64) int {
		cnt := 0
		for _, bus := range time {
			cnt += int(t / int64(bus))
		}
		return cnt
	}

	// 求出解的下界
	for _, bus := range time {
		unitTrips += 1 / float64(bus)
	}
	lowerBound = math.Ceil(float64(totalTrips) / unitTrips)

	// 求出解的上界
	for _, bus := range time {
		upperBound = math.Max(
			upperBound,
			math.Ceil(lowerBound/float64(bus))*float64(bus),
		)
	}

	// 二分搜索
	for {
		var (
			mid        = (lowerBound + upperBound) / 2
			roundedMid = int64(math.Round(mid))
			t0         = trips(roundedMid - 1)
			t1         = trips(roundedMid)
		)
		if t0 < totalTrips && t1 >= totalTrips {
			return roundedMid
		} else if t1 < totalTrips {
			lowerBound = mid
		} else {
			upperBound = mid
		}
	}
}
