class Solution:
    def countAndSay(self, n: int) -> str:
        if n == 1:
            return "1"
        else:
            resultString = ""
            charCount = 0
            prevChar = 0

            for pos, char in enumerate(self.countAndSay(n - 1)):
                if pos == 0:
                    charCount = 1
                    prevChar = char
                elif char == prevChar:
                    charCount += 1
                else:
                    resultString = resultString + str(charCount) + prevChar
                    charCount = 1
                    prevChar = char

            resultString = resultString + str(charCount) + prevChar

            return resultString


sol = Solution()

print(sol.countAndSay(5))
