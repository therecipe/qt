#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QStorageInfo_NewQStorageInfo();
QtObjectPtr QStorageInfo_NewQStorageInfo3(QtObjectPtr dir);
QtObjectPtr QStorageInfo_NewQStorageInfo4(QtObjectPtr other);
QtObjectPtr QStorageInfo_NewQStorageInfo2(char* path);
char* QStorageInfo_DisplayName(QtObjectPtr ptr);
int QStorageInfo_IsReadOnly(QtObjectPtr ptr);
int QStorageInfo_IsReady(QtObjectPtr ptr);
int QStorageInfo_IsRoot(QtObjectPtr ptr);
int QStorageInfo_IsValid(QtObjectPtr ptr);
char* QStorageInfo_Name(QtObjectPtr ptr);
void QStorageInfo_Refresh(QtObjectPtr ptr);
char* QStorageInfo_RootPath(QtObjectPtr ptr);
void QStorageInfo_SetPath(QtObjectPtr ptr, char* path);
void QStorageInfo_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QStorageInfo_DestroyQStorageInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif