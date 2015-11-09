#ifdef __cplusplus
extern "C" {
#endif

void* QSharedMemory_NewQSharedMemory2(void* parent);
void* QSharedMemory_NewQSharedMemory(char* key, void* parent);
int QSharedMemory_Attach(void* ptr, int mode);
void* QSharedMemory_ConstData(void* ptr);
int QSharedMemory_Create(void* ptr, int size, int mode);
void* QSharedMemory_Data(void* ptr);
void* QSharedMemory_Data2(void* ptr);
int QSharedMemory_Detach(void* ptr);
int QSharedMemory_Error(void* ptr);
char* QSharedMemory_ErrorString(void* ptr);
int QSharedMemory_IsAttached(void* ptr);
char* QSharedMemory_Key(void* ptr);
int QSharedMemory_Lock(void* ptr);
char* QSharedMemory_NativeKey(void* ptr);
void QSharedMemory_SetKey(void* ptr, char* key);
void QSharedMemory_SetNativeKey(void* ptr, char* key);
int QSharedMemory_Size(void* ptr);
int QSharedMemory_Unlock(void* ptr);
void QSharedMemory_DestroyQSharedMemory(void* ptr);

#ifdef __cplusplus
}
#endif