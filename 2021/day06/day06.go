package day06

func Run() {
	nums := GetInput()
	Solve(nums, 80)
	Solve(nums, 256)
}

// 587.8 ns/op
func Solve(days [9]int, n int) int {
	for i := 0; i < n; i++ {
		days[0], days[1], days[2], days[3], days[4], days[5], days[6], days[7], days[8] = days[1], days[2], days[3], days[4], days[5], days[6], days[0]+days[7], days[8], days[0]
	}
	return days[0] + days[1] + days[2] + days[3] + days[4] + days[5] + days[6] + days[7] + days[8]
}

func GetInput() [9]int {
	input := []uint8{3, 4, 3, 1, 2, 1, 5, 1, 1, 1, 1, 4, 1, 2, 1, 1, 2, 1, 1, 1, 3, 4, 4, 4, 1, 3, 2, 1, 3, 4, 1, 1, 3, 4, 2, 5, 5, 3, 3, 3, 5, 1, 4, 1, 2, 3, 1, 1, 1, 4, 1, 4, 1, 5, 3, 3, 1, 4, 1, 5, 1, 2, 2, 1, 1, 5, 5, 2, 5, 1, 1, 1, 1, 3, 1, 4, 1, 1, 1, 4, 1, 1, 1, 5, 2, 3, 5, 3, 4, 1, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1, 5, 5, 1, 3, 3, 1, 2, 1, 3, 1, 5, 1, 1, 4, 1, 1, 2, 4, 1, 5, 1, 1, 3, 3, 3, 4, 2, 4, 1, 1, 5, 1, 1, 1, 1, 4, 4, 1, 1, 1, 3, 1, 1, 2, 1, 3, 1, 1, 1, 1, 5, 3, 3, 2, 2, 1, 4, 3, 3, 2, 1, 3, 3, 1, 2, 5, 1, 3, 5, 2, 2, 1, 1, 1, 1, 5, 1, 2, 1, 1, 3, 5, 4, 2, 3, 1, 1, 1, 4, 1, 3, 2, 1, 5, 4, 5, 1, 4, 5, 1, 3, 3, 5, 1, 2, 1, 1, 3, 3, 1, 5, 3, 1, 1, 1, 3, 2, 5, 5, 1, 1, 4, 2, 1, 2, 1, 1, 5, 5, 1, 4, 1, 1, 3, 1, 5, 2, 5, 3, 1, 5, 2, 2, 1, 1, 5, 1, 5, 1, 2, 1, 3, 1, 1, 1, 2, 3, 2, 1, 4, 1, 1, 1, 1, 5, 4, 1, 4, 5, 1, 4, 3, 4, 1, 1, 1, 1, 2, 5, 4, 1, 1, 3, 1, 2, 1, 1, 2, 1, 1, 1, 2, 1, 1, 1, 1, 1, 4}
	var days [9]int
	for _, item := range input {
		days[item]++
	}
	return days
}

// this is slow: 2152 ns/op
// func Solve(days [9]int, n int) int {
//     var sum int
//     var nextDay [9]int
//     for i := 0; i < n; i++ {
//         for day := 8; day >= 0; day-- {
//             if day == 0 {
//                 nextDay[6] += days[day]
//                 nextDay[8] = days[day]
//             } else {
//                 nextDay[day-1] = days[day]
//             }
//         }
//         days = nextDay
//     }
//     for i := 0; i < 9; i++ {
//         sum += days[i]
//     }
//     return sum
// }
