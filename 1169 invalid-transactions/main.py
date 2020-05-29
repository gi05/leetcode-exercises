# MEDIUM
from typing import List


class Solution:
    def invalidTransactions(self, transactions: List[str]) -> List[str]:
        name, amount, time, city = [], [], [], []
        invalid = set()
        keys = list(range(len(transactions)))

        for k, s in enumerate(transactions):
            t = s.split(',')

            name.append(t[0])
            time.append(int(t[1]))
            amount.append(int(t[2]))
            city.append(t[3])

        keys.sort(key=lambda k: time[k])

        for k, v in enumerate(keys):
            if amount[v] > 1000:
                invalid.add(v)

            kk = k-1
            while kk >= 0 and time[keys[k]]-time[keys[kk]] <= 60:
                if name[keys[k]] == name[keys[kk]]:
                    if city[keys[k]] != city[keys[kk]]:
                        invalid.add(keys[k])
                        invalid.add(keys[kk])
                kk -= 1

        return [transactions[i] for i in invalid]


s = Solution()
testCases = [
    ["alice,20,800,mtv", "alice,50,100,beijing"],
    ["alice,20,800,mtv", "alice,50,1200,mtv"],
    ["bob,689,1910,barcelona", "alex,696,122,bangkok", "bob,832,1726,barcelona", "bob,820,596,bangkok", "chalicefy,217,669,barcelona", "bob,175,221,amsterdam"]]

for k, v in enumerate(testCases):
    print(f"Test #{k}")
    print(s.invalidTransactions(v))
    print()
