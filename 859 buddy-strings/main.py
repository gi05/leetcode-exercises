class Solution:
    def buddyStrings(self, a: str, b: str) -> bool:
        if len(a) != len(b):
            return False

        al = list(a)
        bl = list(b)
        d = []
        abc = [0 for _ in range(ord('z')-ord('a'))]

        for k, v in enumerate(al):
            if v != bl[k]:
                d.append(k)
            if len(d) > 2:
                return False
            abc[ord(v) - ord('a')] += 1

        if len(d) == 1:
            return False
        elif len(d) == 0:
            for v in abc:
                if v > 1:
                    return True
            return False

        return al[d[0]] == bl[d[1]] and al[d[1]] == bl[d[0]]


sol = Solution()
# a = "aaaaaaabc"
# b = "aaaaaaacb"

a = "ab"
b = "ab"

# a = "aa"
# b = "aa"

print(sol.buddyStrings(a, b))
