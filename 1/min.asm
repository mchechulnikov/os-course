min:
    movq %rdi, %rax
    cmpq %rsi, %rdi
    jb rdi_ls
    movq %rsi, %rax
rdi_ls:
    ret
