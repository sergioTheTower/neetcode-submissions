class Solution:
    def maxArea(self, heights: List[int]) -> int:
        result = 0
        l, r = 0, len(heights)-1
        while l < r:
            area = (r-l) * min(heights[l], heights[r])
            if result < area:
                result = area
            if heights[l] > heights[r]:
                r-=1
                continue
            l+=1
        return result
        