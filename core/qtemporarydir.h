#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTemporaryDir_NewQTemporaryDir();
QtObjectPtr QTemporaryDir_NewQTemporaryDir2(char* templatePath);
int QTemporaryDir_AutoRemove(QtObjectPtr ptr);
int QTemporaryDir_IsValid(QtObjectPtr ptr);
char* QTemporaryDir_Path(QtObjectPtr ptr);
void QTemporaryDir_SetAutoRemove(QtObjectPtr ptr, int b);
void QTemporaryDir_DestroyQTemporaryDir(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif