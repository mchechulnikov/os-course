pow:
    movq $0, %rbx
    cmp %rsi, %rbx
    je zero_pow
    cmp %rdi, %rbx
    je zero_value
    movq %rdi, %rax
    movq $1, %rbx
mult:
    addq $1, %rbx
    mulq %rdi
    cmp  %rsi, %rbx
    jz exit
    jmp mult
zero_pow:
    movq $1, %rax
    jmp exit
zero_value:
    movq $0, %rax
    jmp exit
exit:
    retq
