#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QFileDevice_AtEnd(QtObjectPtr ptr);
void QFileDevice_Close(QtObjectPtr ptr);
int QFileDevice_Error(QtObjectPtr ptr);
char* QFileDevice_FileName(QtObjectPtr ptr);
int QFileDevice_Flush(QtObjectPtr ptr);
int QFileDevice_Handle(QtObjectPtr ptr);
int QFileDevice_IsSequential(QtObjectPtr ptr);
int QFileDevice_Permissions(QtObjectPtr ptr);
int QFileDevice_SetPermissions(QtObjectPtr ptr, int permissions);
void QFileDevice_UnsetError(QtObjectPtr ptr);
void QFileDevice_DestroyQFileDevice(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif