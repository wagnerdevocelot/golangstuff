#include <cs50.h>
#include <stdio.h>

int main(void){
    int numero = get_int("Digite um numero ai\n");

    if (numero % 2 == 0){
        printf("%i é um numero par\n", numero);
    }
    else {
        printf("%i é um numero impar\n", numero);
    }
    
}

