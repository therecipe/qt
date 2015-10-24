#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSslKey_NewQSslKey();
QtObjectPtr QSslKey_NewQSslKey5(QtObjectPtr other);
void QSslKey_Clear(QtObjectPtr ptr);
int QSslKey_IsNull(QtObjectPtr ptr);
int QSslKey_Length(QtObjectPtr ptr);
void QSslKey_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QSslKey_DestroyQSslKey(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif