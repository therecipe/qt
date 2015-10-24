#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QResource_QResource_RegisterResource(char* rccFileName, char* mapRoot);
int QResource_QResource_UnregisterResource(char* rccFileName, char* mapRoot);
QtObjectPtr QResource_NewQResource(char* file, QtObjectPtr locale);
char* QResource_AbsoluteFilePath(QtObjectPtr ptr);
char* QResource_FileName(QtObjectPtr ptr);
int QResource_IsCompressed(QtObjectPtr ptr);
int QResource_IsValid(QtObjectPtr ptr);
void QResource_SetFileName(QtObjectPtr ptr, char* file);
void QResource_SetLocale(QtObjectPtr ptr, QtObjectPtr locale);
void QResource_DestroyQResource(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif