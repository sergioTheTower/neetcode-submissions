type MinStack struct {
    stack[] int
    min int
}
func Constructor() MinStack {
    return MinStack {
        stack: []int{},
        min: math.MaxInt64,
    }
}
func(this * MinStack) Push(val int) {
    if len(this.stack) == 0 {
        // We only track differences in the stack value.
        this.stack = append(this.stack, 0)
        this.min = val
        return
    }
    this.stack = append(this.stack, val-this.min)
    // If the new value is less than the current min, update min.
    if val < this.min {
        this.min = val
    }
}
func(this * MinStack) Pop() {
    if len(this.stack) == 0 {
        return
    }
    // Get the last value in the stack.
    pop := this.stack[len(this.stack) - 1]
    this.stack = this.stack[:len(this.stack)-1]
    // If the popped value is less than zero.
    if pop < 0 {
        this.min = this.min - pop
    }
}
func(this * MinStack) Top() int {
    top := this.stack[len(this.stack) - 1]
    if top > 0 {
        return top + this.min
    }
    return this.min
}
func(this * MinStack) GetMin() int {
    return this.min
}