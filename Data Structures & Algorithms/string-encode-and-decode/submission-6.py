class Solution:
    # 4#neet4#code4#love3#you
    def encode(self, strs: List[str]) -> str:
        res = ""
        for s in strs:
            l = len(s)
            res += str(l) + "#" + s
        return res



    def decode(self, s: str) -> List[str]:
        l, r = 0, len(s)
        res = []
        while l < r:
            h = s[l:].find('#') + l
            if h == -1:
                break
            num = int(s[l:h])
            word = s[h+1: h+1+num]
            res.append(word)
            l = h+num+1
        return res



            
        
        
