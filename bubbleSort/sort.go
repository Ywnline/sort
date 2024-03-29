package sort

/**
在要排序的切片中，对当前还未排好的序列，从前往后对相邻的两个元素依次进行比较和调整，让较大的数往下沉，较小的往上冒。
即，每当两相邻的元素比较后发现它们的排序与排序要求相反时，就将它们互换。
*/
func bubbleSort(sli []int) []int {


	// 判断是否有排序必要
	length := len(sli)
	if length <= 0 {
		return sli
	}

	//该层循环控制 需要冒泡的轮数
	for i := 0; i < length; i++ {
		//该层循环用来控制每轮 冒出一个数 需要比较的次数
		for j := 0; j < length-1; j++ {
			if sli[i] < sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}

		}

	}

	return sli
}
