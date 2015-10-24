#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSourceLocation_NewQSourceLocation();
QtObjectPtr QSourceLocation_NewQSourceLocation2(QtObjectPtr other);
QtObjectPtr QSourceLocation_NewQSourceLocation3(char* u, int l, int c);
int QSourceLocation_IsNull(QtObjectPtr ptr);
void QSourceLocation_SetUri(QtObjectPtr ptr, char* newUri);
char* QSourceLocation_Uri(QtObjectPtr ptr);
void QSourceLocation_DestroyQSourceLocation(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif