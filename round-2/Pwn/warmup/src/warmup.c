#include <stdio.h>
#include <unistd.h>
#include <stdio.h>
#include <string.h>

void m4gic(){
    FILE *fp = fopen("/home/anonuser/flag.txt", "r");

    if (fp == NULL)
    {
        printf("Error: Please contact @mer4ki.");
        return 1;
    }

    char ch;
    while ((ch = fgetc(fp)) != EOF)
        putchar(ch);

    fclose(fp);

    return 0;
}

int main(){
    setvbuf(stdin, 0, 2, 0);
    setvbuf(stdout, 0, 2, 0);

    volatile int (*fp)();
    char buffer[32];

    puts("Player name: ");
    
    fp = 0;
    
    gets(buffer);
    if(fp) {
      fp();
    }

    puts("Nice name!");
}
