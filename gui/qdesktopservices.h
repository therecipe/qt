#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDesktopServices_QDesktopServices_OpenUrl(char* url);
void QDesktopServices_QDesktopServices_SetUrlHandler(char* scheme, QtObjectPtr receiver, char* method);
void QDesktopServices_QDesktopServices_UnsetUrlHandler(char* scheme);

#ifdef __cplusplus
}
#endif