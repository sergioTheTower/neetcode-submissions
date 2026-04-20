class Solution:
    def isValid(self, s: str) -> bool:
        closeToOpen = { ")" : "(", "]" : "[", "}" : "{" }
        stack = []
        for char in s:
            # Is this an close bracket, check for a match open.
            if char in closeToOpen:
                # Stack not empty and the bracket matches the open bracket.
                if stack and stack[-1] == closeToOpen[char]:
                    stack.pop()
                else:
                    # Imbalance found
                    return False
            else:
                stack.append(char)
        return len(stack) == 0
        