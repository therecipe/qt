#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSharedMemory_NewQSharedMemory2(QtObjectPtr parent);
QtObjectPtr QSharedMemory_NewQSharedMemory(char* key, QtObjectPtr parent);
int QSharedMemory_Attach(QtObjectPtr ptr, int mode);
void QSharedMemory_ConstData(QtObjectPtr ptr);
int QSharedMemory_Create(QtObjectPtr ptr, int size, int mode);
void QSharedMemory_Data(QtObjectPtr ptr);
void QSharedMemory_Data2(QtObjectPtr ptr);
int QSharedMemory_Detach(QtObjectPtr ptr);
int QSharedMemory_Error(QtObjectPtr ptr);
char* QSharedMemory_ErrorString(QtObjectPtr ptr);
int QSharedMemory_IsAttached(QtObjectPtr ptr);
char* QSharedMemory_Key(QtObjectPtr ptr);
int QSharedMemory_Lock(QtObjectPtr ptr);
char* QSharedMemory_NativeKey(QtObjectPtr ptr);
void QSharedMemory_SetKey(QtObjectPtr ptr, char* key);
void QSharedMemory_SetNativeKey(QtObjectPtr ptr, char* key);
int QSharedMemory_Size(QtObjectPtr ptr);
int QSharedMemory_Unlock(QtObjectPtr ptr);
void QSharedMemory_DestroyQSharedMemory(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif