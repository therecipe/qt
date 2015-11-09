#ifdef __cplusplus
extern "C" {
#endif

char* QLibrary_FileName(void* ptr);
int QLibrary_LoadHints(void* ptr);
void QLibrary_SetFileName(void* ptr, char* fileName);
void QLibrary_SetFileNameAndVersion(void* ptr, char* fileName, int versionNumber);
void QLibrary_SetLoadHints(void* ptr, int hints);
void* QLibrary_NewQLibrary(void* parent);
void* QLibrary_NewQLibrary2(char* fileName, void* parent);
void* QLibrary_NewQLibrary4(char* fileName, char* version, void* parent);
void* QLibrary_NewQLibrary3(char* fileName, int verNum, void* parent);
char* QLibrary_ErrorString(void* ptr);
int QLibrary_QLibrary_IsLibrary(char* fileName);
int QLibrary_IsLoaded(void* ptr);
int QLibrary_Load(void* ptr);
void QLibrary_SetFileNameAndVersion2(void* ptr, char* fileName, char* version);
int QLibrary_Unload(void* ptr);
void QLibrary_DestroyQLibrary(void* ptr);

#ifdef __cplusplus
}
#endif