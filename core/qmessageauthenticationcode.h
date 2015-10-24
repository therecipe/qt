#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMessageAuthenticationCode_NewQMessageAuthenticationCode(int method, QtObjectPtr key);
int QMessageAuthenticationCode_AddData2(QtObjectPtr ptr, QtObjectPtr device);
void QMessageAuthenticationCode_AddData3(QtObjectPtr ptr, QtObjectPtr data);
void QMessageAuthenticationCode_AddData(QtObjectPtr ptr, char* data, int length);
void QMessageAuthenticationCode_Reset(QtObjectPtr ptr);
void QMessageAuthenticationCode_SetKey(QtObjectPtr ptr, QtObjectPtr key);
void QMessageAuthenticationCode_DestroyQMessageAuthenticationCode(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif