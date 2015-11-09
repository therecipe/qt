#ifdef __cplusplus
extern "C" {
#endif

int QFileDevice_AtEnd(void* ptr);
void QFileDevice_Close(void* ptr);
int QFileDevice_Error(void* ptr);
char* QFileDevice_FileName(void* ptr);
int QFileDevice_Flush(void* ptr);
int QFileDevice_Handle(void* ptr);
int QFileDevice_IsSequential(void* ptr);
int QFileDevice_Permissions(void* ptr);
int QFileDevice_SetPermissions(void* ptr, int permissions);
void QFileDevice_UnsetError(void* ptr);
void QFileDevice_DestroyQFileDevice(void* ptr);

#ifdef __cplusplus
}
#endif