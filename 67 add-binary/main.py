import math


class Solution:
    def addBinary(self, a: str, b: str) -> str:
        lA = len(a)
        lB = len(b)
        addToNext = 0
        cursor = 0
        result = []

        while (cursor < len(a) and cursor < len(b)) or addToNext > 0:
            if cursor < lA:
                vA = int(a[lA-1-cursor])
            else:
                vA = 0

            if cursor < lB:
                vB = int(b[lB-1-cursor])
            else:
                vB = 0

            result.append(str((vA + vB + addToNext) % 2))
            addToNext = math.floor((vA + vB + addToNext) / 2)

            cursor += 1

        tail = ""
        if lA > cursor:
            tail = a[:lA-cursor]
        elif cursor < lB:
            tail = b[:lB-cursor]

        result.reverse()

        return tail + ''.join(result)


s = Solution()
print(s.addBinary('100010001', '1001'))
