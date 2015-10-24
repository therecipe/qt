#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QValidator_ConnectChanged(QtObjectPtr ptr);
void QValidator_DisconnectChanged(QtObjectPtr ptr);
void QValidator_SetLocale(QtObjectPtr ptr, QtObjectPtr locale);
void QValidator_DestroyQValidator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif