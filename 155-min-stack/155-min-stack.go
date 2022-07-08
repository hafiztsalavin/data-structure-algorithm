// STACK = LIFO

type MinStack struct {
    stack []int
}

func Constructor() MinStack {
    result := &MinStack{stack: []int{}}
    
    return *result
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
}


func (this *MinStack) Pop()  {
    size := len(this.stack)-1
    
    this.stack = this.stack[:size]
}


func (this *MinStack) Top() int {
    top := len(this.stack)-1
    
    return this.stack[top]
}


func (this *MinStack) GetMin() int {   
    min := this.stack[0]
    
    
    for _, value := range this.stack {
        if value < min {
            min = value
        }
    }
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */