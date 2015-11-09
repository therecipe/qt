#ifdef __cplusplus
extern "C" {
#endif

void* QStorageInfo_NewQStorageInfo();
void* QStorageInfo_NewQStorageInfo3(void* dir);
void* QStorageInfo_NewQStorageInfo4(void* other);
void* QStorageInfo_NewQStorageInfo2(char* path);
void* QStorageInfo_Device(void* ptr);
char* QStorageInfo_DisplayName(void* ptr);
void* QStorageInfo_FileSystemType(void* ptr);
int QStorageInfo_IsReadOnly(void* ptr);
int QStorageInfo_IsReady(void* ptr);
int QStorageInfo_IsRoot(void* ptr);
int QStorageInfo_IsValid(void* ptr);
char* QStorageInfo_Name(void* ptr);
void QStorageInfo_Refresh(void* ptr);
char* QStorageInfo_RootPath(void* ptr);
void QStorageInfo_SetPath(void* ptr, char* path);
void QStorageInfo_Swap(void* ptr, void* other);
void QStorageInfo_DestroyQStorageInfo(void* ptr);

#ifdef __cplusplus
}
#endif