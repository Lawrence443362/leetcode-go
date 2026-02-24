package main

import "fmt"

/*
88. Merge Sorted Array

Даны два отсортированных массива nums1 (длина m+n) и nums2 (длина n).
Первые m элементов nums1 — реальные данные, последние n — нули-заглушки.
Слей оба массива в nums1 так, чтобы результат был отсортирован.

Фишка: заполняем nums1 с конца — берём больший из хвостов обоих массивов
и ставим на позицию p. Так мы не затираем непрочитанные элементы.

Пример: nums1 = [1,2,3,0,0,0], m=3, nums2 = [2,5,6], n=3
         → nums1 = [1,2,2,3,5,6]

Паттерн: Two Pointers (с конца)
Время: O(m+n) | Память: O(1)
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	p := m + n -1

	for j >= 0 {
		if i >= 0	 && nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}

		p--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1) // [1 2 2 3 5 6]

	nums2 := []int{0}
	merge(nums2, 0, []int{1}, 1)
	fmt.Println(nums2) // [1]

	nums3 := []int{4, 5, 6, 0, 0, 0}
	merge(nums3, 3, []int{1, 2, 3}, 3)
	fmt.Println(nums3) // [1 2 3 4 5 6]
}
