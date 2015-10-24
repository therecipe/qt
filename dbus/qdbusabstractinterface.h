#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QDBusAbstractInterface_Interface(QtObjectPtr ptr);
int QDBusAbstractInterface_IsValid(QtObjectPtr ptr);
char* QDBusAbstractInterface_Path(QtObjectPtr ptr);
char* QDBusAbstractInterface_Service(QtObjectPtr ptr);
void QDBusAbstractInterface_SetTimeout(QtObjectPtr ptr, int timeout);
int QDBusAbstractInterface_Timeout(QtObjectPtr ptr);
void QDBusAbstractInterface_DestroyQDBusAbstractInterface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif