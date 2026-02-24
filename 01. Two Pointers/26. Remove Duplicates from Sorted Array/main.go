package main

import "fmt"

/*
26. Remove Duplicates from Sorted Array

Дан отсортированный массив целых чисел nums.
Удали дубликаты in-place и верни k — количество уникальных элементов.
Первые k элементов массива должны содержать уникальные значения в исходном порядке.
Что стоит после них — не важно. Новый массив выделять нельзя.

Пример: [1,1,1,1,3,4,5,5,6,7,7,8,9,9,10]
         → k=9, nums[:9] = [1,3,4,5,6,7,8,9,10]

По сути — массив пришёл на вечеринку, а охранник пускает только по одному от каждого числа.
Остальные стоят в конце и грустят, но их никто не проверяет.

Паттерн: Two Pointers (slow/fast)
Время: O(n) | Память: O(1)
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0;
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++;
			nums[slow] = nums[fast];
		}
	}

	return slow + 1
}

func main() {
	nums := []int{1, 1, 1, 1, 3, 4, 5, 5, 6, 7, 7, 8, 9, 9, 10}
	k := removeDuplicates(nums)
	fmt.Println("k =", k)
	fmt.Println("nums[:k] =", nums[:k])
}