class Solution:
    def defangIPaddr(self, address: str) -> str:
        res = [c if c != '.' else '[.]' for c in address]

        return ''.join(res)


sol = Solution()
address = "1.1.1.1"
address = "255.100.50.0"

print(sol.defangIPaddr(address))
