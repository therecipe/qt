#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSslError_NewQSslError();
QtObjectPtr QSslError_NewQSslError2(int error);
QtObjectPtr QSslError_NewQSslError3(int error, QtObjectPtr certificate);
QtObjectPtr QSslError_NewQSslError4(QtObjectPtr other);
int QSslError_Error(QtObjectPtr ptr);
char* QSslError_ErrorString(QtObjectPtr ptr);
void QSslError_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QSslError_DestroyQSslError(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif