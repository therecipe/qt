#ifdef __cplusplus
extern "C" {
#endif

void* QTemporaryFile_NewQTemporaryFile();
void* QTemporaryFile_NewQTemporaryFile3(void* parent);
void* QTemporaryFile_NewQTemporaryFile2(char* templateName);
void* QTemporaryFile_NewQTemporaryFile4(char* templateName, void* parent);
int QTemporaryFile_AutoRemove(void* ptr);
void* QTemporaryFile_QTemporaryFile_CreateNativeFile(void* file);
void* QTemporaryFile_QTemporaryFile_CreateNativeFile2(char* fileName);
char* QTemporaryFile_FileName(void* ptr);
char* QTemporaryFile_FileTemplate(void* ptr);
int QTemporaryFile_Open(void* ptr);
void QTemporaryFile_SetAutoRemove(void* ptr, int b);
void QTemporaryFile_SetFileTemplate(void* ptr, char* name);
void QTemporaryFile_DestroyQTemporaryFile(void* ptr);

#ifdef __cplusplus
}
#endif