class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        lookup = {}
        for char in s:
            lookup[char] = lookup.get(char, 0) + 1  # Initialize or increment count
        
        for char in t:
            if char not in lookup or lookup[char] == 0:
                return False
            lookup[char] -= 1  # Decrement count

        # Ensure all counts are zero
        for value in lookup.values():
            if value != 0:
                return False
        
        return True
