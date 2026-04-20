class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res = [0] * len(temperatures)
        stack = []
        for i, t in enumerate(temperatures):
            print(i,t)
            # Stack is not empty AND new temp is higher than past temp.
            while stack and t > stack[-1][0]:
                print("Not Empty/New Temp")
                stackT, prev_index = stack.pop()
                # Subtrack the two indexs to determine the diff/number of warmer days.
                res[prev_index] = i - prev_index
                print("*Answer: ",prev_index, i-prev_index)
            stack.append((t,i))
            print(stack)
        return res
        