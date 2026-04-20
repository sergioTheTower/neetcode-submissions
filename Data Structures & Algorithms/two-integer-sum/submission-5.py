class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        diff_tracker = {}
        for idx, n in enumerate(nums):
            diff = target - n
            if diff in diff_tracker: 
                return [diff_tracker[diff], idx]
            diff_tracker[n] = idx
        return []