class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        for n in nums:
            #  Get the count of all the nums.
            count[n] = 1 + count.get(n, 0)

        buckets = [[] for _ in range(len(nums) + 1)]
        print("Buckets: ", buckets)
        for num, cnt in count.items():
            # Index represents the count in buckets, and the numbers that have that count.
            buckets[cnt].append(num)

        result = []
        for i in range(len(buckets) - 1, 0, -1):
            for n in buckets[i]:
                result.append(n)
                if len(result) == k:
                    return result

        return result
        