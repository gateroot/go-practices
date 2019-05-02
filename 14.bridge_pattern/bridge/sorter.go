package bridge

type Sorter struct {
	SortImpl SortImpl
}

func (sorter *Sorter) Sort(values []int) []int {
	return sorter.SortImpl.Sort(values)
}

type SortImpl interface {
	Sort([]int) []int
}

type QuickSortImpl struct{}

func (quick *QuickSortImpl) Sort(values []int) (ret []int) {
	// 要素数が 1 以下の配列はそれ以上細分化してソートする必要がない
	if len(values) < 2 {
		return values
	}

	// 配列の先頭をピボット (基準値) に選ぶ
	pivot := values[0]

	// ピボットを基準にして値の大小で配列をふたつに分割する
	left := []int{}
	right := []int{}
	for _, v := range values[1:] {
		if v > pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	// 分割した配列をそれぞれ再帰的にソートする
	left = quick.Sort(left)
	right = quick.Sort(right)

	/*
	   left(小さい) + pivot(基準値) + right(大きい) の順番で配列を組み立てる
	   ここが実際のソート処理
	*/
	ret = append(left, pivot)
	ret = append(ret, right...)

	return
}

type BubbleSortImple struct{}

func (bubble *BubbleSortImple) Sort(values []int) []int {
	for i, _ := range values {
		for j := i; j < len(values); j++ {
			if i > j {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}
