from typing import List


class Solution:
    def luckyNumbers(self, matrix: List[List[int]]) -> List[int]:

        minV = []
        maxV = []
        res = []

        for k in range(len(matrix)):
            minV.append(min(matrix[k]))

        for k in range(len(matrix[0])):
            tmpL = [matrix[rk][k] for rk in range(len(matrix))]
            maxV.append(max(tmpL))

        for rk, rv in enumerate(matrix):
            for ck, cv in enumerate(rv):
                if cv == minV[rk] and cv == maxV[ck]:
                    res.append(cv)

        return res


sol = Solution()

# m = [[3, 7, 8], [9, 11, 13], [15, 15, 17]]
# m = [[1, 10, 4, 2], [9, 3, 8, 7], [12, 16, 17, 12]]
m = [[1]]
print(sol.luckyNumbers(m))
