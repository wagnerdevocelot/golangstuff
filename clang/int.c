#include <cs50.h>
#include <stdio.h>

int main(void){
    int age = get_int("Qual sua idade?\n");
    printf("Você está pelo menos %i dias mais velho\n", age * 365);
}


