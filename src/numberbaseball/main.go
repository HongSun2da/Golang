package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BallResult struct {
	strikes int
	balls   int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 무작위 숫자 3개 만들기
	numbers := MakeNumbers()

	cnt := 0

	for {
		// 사용자의 입력을 받는다.
		inputNumbers := InputNumbers()

		// 결과를 비교 한다.
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 결과 확인
		if IsGameEnd(result) {
			// 게임끝
			break
		}
	}

	// 몇번만에
	fmt.Printf("%d 번만에 맞췄습니다.\n", cnt)
}

func MakeNumbers() [3]int {
	// 0~9 사이의 무작위 숫자 3개를 반환한다.
	var rst [3]int

	// rand.Intn(10)
	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			fmt.Println(n)
			duplicated := false

			for j := 0; j < i; j++ {
				if rst[j] == n {
					// 중복 발생, 다시 뽑기
					duplicated = true
					break
				}
			}

			if !duplicated {
				rst[i] = n
				break
			}
		}
	}

	fmt.Println(rst)

	return rst
}

func InputNumbers() [3]int {
	// 키보드로 부터 0~9 사이의 무작위 숫자 3개를 반환한다.
	var rst [3]int

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력 하세요.")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n {
					// 중복 발생, 다시 뽑기
					fmt.Println("숫자가 겹치지 않아야 합니다.")
					duplicated = true
					break
				}
			}

			if duplicated {
				success = false
				break
			}

			rst[idx] = n
			idx++
		}
		if !success {
			continue
		}
		break
	}

	rst[0], rst[2] = rst[2], rst[0]

	fmt.Println(rst)
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) BallResult {
	// 두개의 3개 숫자를 비교해서 결과 반환
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}

	return BallResult{strikes, balls}
}

func PrintResult(result BallResult) {
	fmt.Printf("(%dS, %dB)", result.strikes, result.balls)
}

func IsGameEnd(result BallResult) bool {
	// 비교 결과가 3 스트라이크 확인
	return result.strikes == 3
}
