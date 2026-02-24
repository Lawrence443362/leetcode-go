package main

import "fmt"

/*
66. Plus One

Дан массив цифр, представляющий целое число (digits[0] — старший разряд).
Прибавь 1 и верни результат как массив цифр.

По сути — ты кассир, и тебе сказали "добавь копейку". Если в разряде 9 — он
превращается в 0 и ты несёшь единицу дальше. Если дошёл до конца и всё ещё
несёшь — значит было 999...9 и теперь стало 1000...0.

Пример: [1,2,9] → [1,3,0]
         [9,9,9] → [1,0,0,0]

Паттерн: Проход справа налево
Время: O(n) | Память: O(1) (кроме случая 999→1000)
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1

			return digits
		} else {
			digits[i] = 0
		}
	}

	digits = make([]int, len(digits) + 1);
	digits[0] = 1
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 9}))    // [1 3 0]
	fmt.Println(plusOne([]int{9, 9, 9}))    // [1 0 0 0]
	fmt.Println(plusOne([]int{1, 2, 3}))    // [1 2 4]
	fmt.Println(plusOne([]int{0}))           // [1]
}
