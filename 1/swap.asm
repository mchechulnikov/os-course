swap:
    pushq (%RSI)
    pushq (%RDI)
    popq (%RSI)
    popq (%RDI)
    retq
