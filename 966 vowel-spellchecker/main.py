# MEDIUM
import re
from typing import List


class Solution:
    def spellchecker(self, wordlist: List[str], queries: List[str]) -> List[str]:
        reVowel = re.compile('[aeiou]')
        wordExact = {}
        wordLower = {}
        wordVowels = {}
        result = []

        for w in wordlist:
            wordExact[w] = w

            wLower = w.lower()
            if not wordLower.get(wLower):
                wordLower[wLower] = w

            wVowel = reVowel.sub('_', wLower)
            if not wordVowels.get(wVowel):
                wordVowels[wVowel] = w

        for q in queries:
            qLower = q.lower()
            qVowel = reVowel.sub('_', qLower)

            if wordExact.get(q):
                result.append(q)
            elif wordLower.get(qLower):
                result.append(wordLower[qLower])
            elif wordVowels.get(qVowel):
                result.append(wordVowels[qVowel])
            else:
                result.append("")

        return result


s = Solution()

testCases = [
    (["KiTe", "kite", "hare", "Hare"], ["kite", "Kite", "KiTe",
                                        "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"])
]

for k, v in enumerate(testCases):
    print(f"Test #{k}")
    print(s.spellchecker(v[0], v[1]))
    print()
