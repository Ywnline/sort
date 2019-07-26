package sort

import "sort"

func QuickSort(sli []int) []int {

	// 判断是否有必要进行排序
	length := len(sli)
	if length <= 1 {
		return sli
	}

	// 取出切片中的第一个元素基数
	base_num := sli[0]

	//初始化左右两个切片
	left_num := []int{}  //小于基准的
	right_num := []int{} //小于基准的

	//遍历除了标尺外的所有元素，按照大小关系放入左右两个切片内
	for i := 1; i < length; i++ {

		if base_num > sli[i] {
			//放入左边切片
			left_num = append(left_num, sli[i])
		} else {
			//放入右边切片
			right_num = append(right_num, sli[i])
		}
	}

	//再分别对左边和右边的切片进行相同的排序处理方式递归调用这个函数
	left_num = QuickSort(left_num)
	right_num = QuickSort(right_num)

	// 合并
	left_num = append(left_num, base_num)

	return append(left_num, right_num...)
}

/**
go 切片排序
*/
func goSort(sli []int) []int {
	// go 语言支持
	sortLi := sort.IntSlice(sli)

	//sortLi.Sort()
	sort.Sort(sortLi)

	return sortLi
}
