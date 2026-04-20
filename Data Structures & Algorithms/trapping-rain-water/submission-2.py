class Solution:
    def trap(self, height: List[int]) -> int:
        left_pointer, right_pointer = 0, len(height)-1
        left_max, right_max = height[left_pointer], height[right_pointer]
        result = 0
        while left_pointer < right_pointer:
            if left_max < right_max:
                left_pointer +=1
                left_max = max(left_max, height[left_pointer])
                result += left_max - height[left_pointer]
            else:
                right_pointer -=1
                right_max = max(right_max, height[right_pointer])
                result += right_max - height[right_pointer]
        return result

        