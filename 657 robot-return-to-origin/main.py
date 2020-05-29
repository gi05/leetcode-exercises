class Solution:
    def judgeCircle(self, moves: str) -> bool:
        hor, ver = 0, 0
        for v in moves:
            if v == 'L':
                hor += 1
            elif v == 'R':
                hor -= 1
            elif v == 'U':
                ver += 1
            elif v == 'D':
                ver -= 1
        return hor == 0 and ver == 0
