#include <cs50.h>
#include <stdio.h>


int main(void){
    string answer = get_string("Qual seu nome?\n");
    printf("Hello, %s", answer);
}