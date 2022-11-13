// hello.c
#include "hello.h"


void hello() {
    printf("hello %s\n", "hoge");
}
void helloint(int val) {
    printf("hello %d\n", val);
}
void hellostr(const char *val) {
    printf("hello %s\n", val);
}
