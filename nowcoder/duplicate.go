package nowcoder

// nowcoder title: 数组中重复的数字
// url: https://www.nowcoder.com/practice/6fe361ede7e54db1b84adc81d09d8524?tpId=13&tqId=11203&tab=answerKey&from=cyc_github
/*
解题思路： 
	长为 n 的数组中的任意数字 i 都是 0<=i<n 的，所以我们只需要将每个数字移回自己的位置，即保证 numbers[i] == i，
	需要做的操作是 numbers[i], numbers[numbers[i]] = numbers[numbers[i]],numbers[i]，
	如果数组中有重复的数字，在移动的过程中必然会出现 numbers[i] == numbers[numbers[i]] 的情况
*/
func Duplicate( numbers []int ) int {
    var result = -1
    for i:=0;i<len(numbers);i++{
        for numbers[i]!=i{
            if numbers[i] == numbers[numbers[i]]{
            return numbers[i]
        }
            numbers[i], numbers[numbers[i]] = numbers[numbers[i]],numbers[i]
        }
    }
    return result
}