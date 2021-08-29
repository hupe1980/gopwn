#include <stdio.h>
#include <string.h>

int main() {
    char pass[10];
    printf("Input the password:\n");
    fflush(stdout);
    scanf("%s", pass);

    if (strcmp(pass, "TopS3cr3t!") == 0) {
        printf("Correct password\n");
        return 0;
    }

    printf("Wrong password, try another\n");
    return 0;
}
