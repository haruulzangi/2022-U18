#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

const char *map_locations[] = {"Santa Fe, New Mexico, USA", 
    "Bermuda Triangle, Atlantic Ocean", 
    "Bennett Island, Russia",
    "Rabat, Morocco",
    "Thimphu, Bhutan",
    "Tokyo, Japan",
    "Venice, Italy",
    "Venice, Italy",
    NULL
    };

int getFlag(){
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

int mapFunc(){
    char buffer[128];

    while(buffer[0]!='9'){
        banner();
        printf("># ");
        fgets(buffer, 128, stdin);
        if (buffer[0] < '7' && buffer[0] >= '0') {
            printf("%s\n\n", map_locations[atoi(buffer)]);
        }
        else {
            printf("Оруулсан утга буруу байна, оруулсан утга: ");
            printf(buffer);
            exit(1);
        }
    }
}

void asciiart(){
    printf(
    "⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿ \n"
    "⠀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠏⠀⠀\n"
    "⠀⠀⠀⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⠀⠀⠀\n"
    "⠀⠀⠀⠀⠀⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠀⠀⠀⠀⠀\n"
    "⠀⠀⠀⠀⠀⠀⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⠀⠀⠀⠀⠀⠀\n"
    "⠀⠀⠀⠀⠀⠀⠀⠀⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠋⠀⠀⠀⠀⠀⠀⠀⠀\n"
    "⡄⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⣰\n"
    "⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠻⢿⣿⣿⣿⣿⣿⠋⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣾\n"
    "⣿⣿⣷⡄⠀⠀⠀⠀⠀⠀⠀⠀⠘⢿⣿⣿⣿⣥⣤⣤⣤⣤⣤⣤⣤⣤⣴⣿⣿⣿\n"
    "⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠛⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\n"
    "⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\n"
    "⣿⣿⣿⣿⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\n"
    "⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⢈⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\n");
}

void banner(){
    printf(
    "Та Valorant тоглоомын map-г сонгоод байршлийг нь хараарай.\n\n"
    "[0] Fracture\n"    
    "[1] Breeze\n"	
    "[2] Icebox\n" 	
    "[3] Bind\n"		
    "[4] Haven\n"		
    "[5] Split\n"		
    "[6] Ascent\n"  
    "[7] The Range\n"
    "[9] Гарах\n"
    );
}


int main(int argc, char **argv){
    setvbuf(stdin, 0, 2, 0);
    setvbuf(stdout, 0, 2, 0);

    asciiart();

    mapFunc();

    return 0;
}
