
MEMORY
{
    FLASH_TEXT (rw) : ORIGIN = 0x00000000+0x0000, LENGTH = 0x00080000-0x0000
    RAM (xrw)       : ORIGIN = 0x20000000, LENGTH = 0x00030000
}

_stack_size = 4K;

INCLUDE "targets/arm.ld"
