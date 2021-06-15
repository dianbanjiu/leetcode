package nowcoder

// nowcode title: 二维数组中的查找
/* 
解题思路： 
	根据数组的排列顺序，可以从数组的左上角开始，依次进行下面三个判断
	1. 进行下面每个步骤之前都需要判断当前元素是否等于 target，等于直接返回 true
	2. 当前元素是否大于目标值，大于的话将列指针向前移动一位
	3. 重复步骤 2，直到当前元素的值小于 target
	4. 当前元素是否小于目标值，小于的话将行指针下移
	5. 重复步骤 4，直到循环结束，此时返回 false

*/
func Find( target int ,  array [][]int ) bool {
    // write code here
    i,j := 0,len(array[0])-1
    for ;i<len(array)&&j>=0;{
        if array[i][j] == target{
            return true
        }else if array[i][j] > target{
            j--
        }else if array[i][j] < target{
            i++
        }
    }
    return false
}