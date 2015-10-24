#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTemporaryFile_NewQTemporaryFile();
QtObjectPtr QTemporaryFile_NewQTemporaryFile3(QtObjectPtr parent);
QtObjectPtr QTemporaryFile_NewQTemporaryFile2(char* templateName);
QtObjectPtr QTemporaryFile_NewQTemporaryFile4(char* templateName, QtObjectPtr parent);
int QTemporaryFile_AutoRemove(QtObjectPtr ptr);
QtObjectPtr QTemporaryFile_QTemporaryFile_CreateNativeFile(QtObjectPtr file);
QtObjectPtr QTemporaryFile_QTemporaryFile_CreateNativeFile2(char* fileName);
char* QTemporaryFile_FileName(QtObjectPtr ptr);
char* QTemporaryFile_FileTemplate(QtObjectPtr ptr);
int QTemporaryFile_Open(QtObjectPtr ptr);
void QTemporaryFile_SetAutoRemove(QtObjectPtr ptr, int b);
void QTemporaryFile_SetFileTemplate(QtObjectPtr ptr, char* name);
void QTemporaryFile_DestroyQTemporaryFile(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif