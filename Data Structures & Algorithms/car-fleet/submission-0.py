class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        # Join position and speed.
        pair = [[p,s] for p,s in zip(position, speed)]
        stack = []
        # Loop over the pairs in reverse order.
        for p,s in sorted(pair)[::-1]:
            stack.append((target-p)/s)
            # There are at least two cars and the end car and first car collide.
            if len(stack) >= 2 and stack[-1] <= stack[-2]:
                stack.pop()
        return len(stack)