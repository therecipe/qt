#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QLibrary_FileName(QtObjectPtr ptr);
int QLibrary_LoadHints(QtObjectPtr ptr);
void QLibrary_SetFileName(QtObjectPtr ptr, char* fileName);
void QLibrary_SetFileNameAndVersion(QtObjectPtr ptr, char* fileName, int versionNumber);
void QLibrary_SetLoadHints(QtObjectPtr ptr, int hints);
QtObjectPtr QLibrary_NewQLibrary(QtObjectPtr parent);
QtObjectPtr QLibrary_NewQLibrary2(char* fileName, QtObjectPtr parent);
QtObjectPtr QLibrary_NewQLibrary4(char* fileName, char* version, QtObjectPtr parent);
QtObjectPtr QLibrary_NewQLibrary3(char* fileName, int verNum, QtObjectPtr parent);
char* QLibrary_ErrorString(QtObjectPtr ptr);
int QLibrary_QLibrary_IsLibrary(char* fileName);
int QLibrary_IsLoaded(QtObjectPtr ptr);
int QLibrary_Load(QtObjectPtr ptr);
void QLibrary_SetFileNameAndVersion2(QtObjectPtr ptr, char* fileName, char* version);
int QLibrary_Unload(QtObjectPtr ptr);
void QLibrary_DestroyQLibrary(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif