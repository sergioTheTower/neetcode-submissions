func maxArea(heights []int) int {
    var result int
    left, right := 0, len(heights)-1
    for left < right {
        area := (right-left) * minValue(heights[left], heights[right])
        if area > result {
            result = area
        }
        if heights[left] > heights[right] {
            right-=1
            continue
        }
        left +=1
    }
return result
}


func minValue(x,y int) int {
    if x > y {
        return y
    }
    return x
}