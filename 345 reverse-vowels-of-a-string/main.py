class Solution:
    def reverseVowels(self, s: str) -> str:
        res = list(s)
        sample = "aeiouAEIOU"
        i1 = 0
        i2 = len(s)-1
        while i1 <= i2:
            if s[i1] in sample:
                while not s[i2] in sample:
                    i2 -= 1
                if i1 < i2:
                    res[i1] = s[i2]
                    res[i2] = s[i1]
                i2 -= 1
            i1 += 1

        return "".join(res)


sol = Solution()
s = "hello"
# s = "leetcode"

print(sol.reverseVowels(s))
