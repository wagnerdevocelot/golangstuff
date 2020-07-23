#include <stdio.h>
#include <cs50.h>


int main(void){
    float valor = get_float("Qual o preço?\n");
    printf("O valor total é de %.2f Reais\n", valor * 1.0625 );
}