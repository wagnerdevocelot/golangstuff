#include <stdio.h>
#include <cs50.h>

int main()
{
int idade;
char nome[30];
printf("Digite sua Idade: ");
scanf("%d",&idade);
printf("Seu Nome: ");
scanf("%s",nome);    /* Strings não utilizar '&' na leitura */
printf("%s Sua idade e' %d anos. \n", nome, idade);
}