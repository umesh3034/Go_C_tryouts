package main

import "fmt"

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	var bribecount, i, j, arrlen int32
	arrlen = int32(len(q))
	for i = arrlen - 1; i >= 0; i-- {
		/* To figure too chaotic case */
		if q[i]-int32(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		/* to check bribecount, start from end
		 * and check how may person got ahead
		 * of you, see below condition
		 */
		for j = i - 1; i > 0 && j >= 0; j-- {
			if q[j] > q[i] { //overtaken
				bribecount++
			}
		}
	}
	fmt.Println(bribecount)
}

func main() {
	minimumBribes([]int32{5, 1, 2, 3, 7, 8, 6, 4})
	minimumBribes([]int32{1, 2, 5, 3, 7, 8, 6, 4})
}
