package sort

/**
在要排序的一切片中，假设前面的元素已经是排好顺序的，现在要把第n个元素插到前面的有序切片中，使得这n个元素也是排好顺序的。
如此反复循环，直到全部排好顺序。
*/
func insertSort(sli []int) []int {
	length := len(sli)

	if length <= 0 {
		return sli
	}

	for i := 0; i < length; i++ {
		tmp := sli[i]
		//内层循环控制，比较并插入
		for j := i - 1; j >= 0; j-- {

			if sli[j] > tmp {
				//发现插入的元素要小，交换位置，将后边的元素与前面的元素互换
				sli[j+1], sli[j] = sli[j], tmp
			} else {
				//如果碰到不需要移动的元素，则前面的就不需要再次比较了。
				break
			}
		}
	}
	return sli
}
