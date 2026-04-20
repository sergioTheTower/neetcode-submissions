class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        count_dict = {}
        for char in s:
            count_dict[char] = count_dict.get(char, 0) + 1
        for char in t:
            if char not in count_dict or count_dict[char] == 0:
                return False
            count_dict[char] -= 1
        return True