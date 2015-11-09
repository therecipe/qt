#ifdef __cplusplus
extern "C" {
#endif

void* QSaveFile_NewQSaveFile2(void* parent);
void* QSaveFile_NewQSaveFile(char* name);
void* QSaveFile_NewQSaveFile3(char* name, void* parent);
void QSaveFile_CancelWriting(void* ptr);
int QSaveFile_Commit(void* ptr);
int QSaveFile_DirectWriteFallback(void* ptr);
char* QSaveFile_FileName(void* ptr);
int QSaveFile_Open(void* ptr, int mode);
void QSaveFile_SetDirectWriteFallback(void* ptr, int enabled);
void QSaveFile_SetFileName(void* ptr, char* name);
void QSaveFile_DestroyQSaveFile(void* ptr);

#ifdef __cplusplus
}
#endif