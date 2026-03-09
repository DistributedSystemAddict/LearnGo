#include "textflag.h"

TEXT ·Sum(SB), NOSPLIT, $0-32
    MOVQ arr_base+0(FP), SI
    MOVQ arr_len+8(FP), CX

    XORQ AX, AX      // sum = 0
    XORQ DX, DX      // i = 0

loop:
    CMPQ DX, CX
    JGE done

    MOVQ (SI)(DX*8), BX
    ADDQ BX, AX

    INCQ DX
    JMP loop

done:
    MOVQ AX, ret+24(FP)
    RET
