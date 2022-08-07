type NumArray struct {
    ArrayNum []int
}


func Constructor(nums []int) NumArray {
    return NumArray{
        ArrayNum : nums,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    result := 0
    for i := left; i <= right; i++{
        result = result + this.ArrayNum[i]
    } 
    
    return result
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */