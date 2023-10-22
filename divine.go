package main

var divineTable = map[int]string{
	1: "大安：诸事顺利。",
	2: "留连：运气平平。",
	3: "速喜：时机已到。",
	4: "赤口：谨防小人。",
	5: "小吉：好事发生，耐心等待。",
	6: "空亡：诸事不顺利，事事小心。",
}

func TimeTable(hour int) int {
	switch {
	case hour >= 23 || hour < 1:
		return 1
	case hour >= 1 && hour < 3:
		return 2
	case hour >= 3 && hour < 5:
		return 3
	case hour >= 5 && hour < 7:
		return 4
	case hour >= 7 && hour < 9:
		return 5
	case hour >= 9 && hour < 11:
		return 6
	case hour >= 11 && hour < 13:
		return 7
	case hour >= 13 && hour < 15:
		return 8
	case hour >= 15 && hour < 17:
		return 9
	case hour >= 17 && hour < 19:
		return 10
	case hour >= 19 && hour < 21:
		return 11
	case hour >= 21 && hour < 23:
		return 12
	default:
		return 0
	}
}
