#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSaveFile_NewQSaveFile2(QtObjectPtr parent);
QtObjectPtr QSaveFile_NewQSaveFile(char* name);
QtObjectPtr QSaveFile_NewQSaveFile3(char* name, QtObjectPtr parent);
void QSaveFile_CancelWriting(QtObjectPtr ptr);
int QSaveFile_Commit(QtObjectPtr ptr);
int QSaveFile_DirectWriteFallback(QtObjectPtr ptr);
char* QSaveFile_FileName(QtObjectPtr ptr);
int QSaveFile_Open(QtObjectPtr ptr, int mode);
void QSaveFile_SetDirectWriteFallback(QtObjectPtr ptr, int enabled);
void QSaveFile_SetFileName(QtObjectPtr ptr, char* name);
void QSaveFile_DestroyQSaveFile(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif