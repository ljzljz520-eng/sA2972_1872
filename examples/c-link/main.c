#include <stdio.h>
#include "libparking_capi.h"

int main(void) {
    char *credential = ParkingCredentialCreate("沪A12345", 1748765700LL, "B2");
    char *status = ParkingCredentialValidate(credential, 1748774700LL, "B2");
    if (credential == NULL || status == NULL) {
        return 1;
    }
    printf("credential=%s\nstatus=%s\n", credential, status);
    ParkingStringFree(status);
    ParkingStringFree(credential);
    return 0;
}
