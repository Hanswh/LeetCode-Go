package leetcode

func slowestKey(releaseTimes []int, keysPressed string) byte {
	res := keysPressed[0]
	maxTime := releaseTimes[0]

	for i := 1; i < len(keysPressed); i++ {
		pressedTime := releaseTimes[i] - releaseTimes[i-1]
		if pressedTime > maxTime {
			maxTime = pressedTime
			res = keysPressed[i]
		} else if pressedTime == maxTime && res < keysPressed[i] {
			res = keysPressed[i]
		}
	}

	return res
}
