# MEDIUM
from typing import List
import math


class Solution:
    def nextGreaterElement(self, n: int) -> int:
        digits = []
        result = 0

        t = n
        while t > 0:
            digits.append(t % 10)
            t = math.floor(t / 10)

        k = 1
        while k < len(digits) and digits[k] >= digits[k-1]:
            k += 1

        if k == len(digits):
            return -1

        kk = 0
        while digits[kk] <= digits[k]:
            kk += 1

        digits[k], digits[kk] = digits[kk], digits[k]

        digits = sorted(digits[:k], reverse=True) + digits[k:]
        for d in reversed(digits):
            result = result * 10 + d
            if result > 2 << 30:
                return -1

        return result


s = Solution()

testCases = [
    11,
    1999999999,
    2147483647,
    1234,
    4321,
    214321]

for k, v in enumerate(testCases):
    print(f"Test #{k}")
    print(s.nextGreaterElement(v))
    print()
