#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFileIconProvider_NewQFileIconProvider();
int QFileIconProvider_Options(QtObjectPtr ptr);
void QFileIconProvider_SetOptions(QtObjectPtr ptr, int options);
char* QFileIconProvider_Type(QtObjectPtr ptr, QtObjectPtr info);
void QFileIconProvider_DestroyQFileIconProvider(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif