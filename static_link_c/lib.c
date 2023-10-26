#include <stdio.h>
int Add(int a, int b)
{
    printf("Welcome from external C function\n");
    return a + b;
}

// gcc -shared -fpic lib.c -o ./libadd.a
