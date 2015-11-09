#ifdef __cplusplus
extern "C" {
#endif

void* QTemporaryDir_NewQTemporaryDir();
void* QTemporaryDir_NewQTemporaryDir2(char* templatePath);
int QTemporaryDir_AutoRemove(void* ptr);
int QTemporaryDir_IsValid(void* ptr);
char* QTemporaryDir_Path(void* ptr);
void QTemporaryDir_SetAutoRemove(void* ptr, int b);
void QTemporaryDir_DestroyQTemporaryDir(void* ptr);

#ifdef __cplusplus
}
#endif