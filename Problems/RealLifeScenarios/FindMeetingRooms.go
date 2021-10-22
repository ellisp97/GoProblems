package main

import "sort"

func minMeetingRooms(meetingTimes [][]int) int {
	if len(meetingTimes) == 0 {
		return 0
	}

	sort.SliceStable(meetingTimes, func(i, j int) bool {
		return meetingTimes[i][1] < meetingTimes[j][1]
	})

	pq := make(PriorityQueue, 5)
}
