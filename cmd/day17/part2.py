def do(inp: list[int], cur: str) -> str:
    if not inp:
        return cur
    num = inp.pop()
    for q in ['000', '001', '010', '011', '100', '101', '110', '111']:
        mayb = cur + q
        # A%8 XOR 5 XOR 6 XOR A/(2^(A%8 XOR 5))
        e = int(mayb, 2)
        w = (e % 8) ^ 5
        r = ((e % 8) ^ 5 ^ 6 ^ int(e / (2**w))) % 8
        if r == num and (res := do(inp.copy(), mayb)):
            return res
    return ""


instr = [2,4,1,5,7,5,1,6,0,3,4,6,5,5,3,0]
res = do(instr, "")
print(res)
print(int(res, 2))