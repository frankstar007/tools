/**
 * created by
 * @project tools
 * @author frankstar
 * @date 2023/2/14
 * @contact frankstarye@tencent.com
 **/

package collection

type Element interface {}

//Int64ArraySplit 先实现int64的数组分组 后续泛型实现
func Int64ArraySplit(arr []int64, num int64) [][]int64 {
	size := int64(len(arr))

	if size <= num {
		return [][]int64{arr}
	}

	groups := size / num
	if size % num != 0{
		groups += 1
	}

	var results [][]int64
	var start, end ,index int64
	for index = 1; index <= groups; index ++ {
		end = index * num
		if index != groups {
			results = append(results, arr[start:end])
		} else {
			results = append(results, arr[start:])
		}
		start = index * num
	}

	return results
}
