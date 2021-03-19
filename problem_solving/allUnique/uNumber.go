package main

// func FindUniq(arr []float32) float32 {
// 	keys := make(map[float32]bool)
// 	var list float32
// 	for _, entry := range arr {
// 		if _, value := keys[entry]; !value {
// 			keys[entry] = true
// 			list = entry
// 		}
// 	}
// 	fmt.Println(list)
// 	return list

// }

// func main() {
// 	var arr []float32
// 	arr = []float32{0, 0, 0.55, 0, 0}
// 	FindUniq(arr)
// }

func FindUniq(arr []int) bool {
	keys := make(map[int]bool)

	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			return true
		}
	}

	return false

}

func main() {
	var arr []int
	arr = []int{0, 0, 2, 0, 0}
	FindUniq(arr)
}
