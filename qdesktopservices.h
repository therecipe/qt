#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Static Public Members
int QDesktopServices_OpenUrl_String(char* url);
void QDesktopServices_SetUrlHandler_String_QObject_String(char* scheme, QtObjectPtr receiver, char* method);
void QDesktopServices_UnsetUrlHandler_String(char* scheme);

#ifdef __cplusplus
}
#endif
