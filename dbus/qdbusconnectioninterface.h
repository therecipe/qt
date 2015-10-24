#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QDBusConnectionInterface_ConnectServiceRegistered(QtObjectPtr ptr);
void QDBusConnectionInterface_DisconnectServiceRegistered(QtObjectPtr ptr);
void QDBusConnectionInterface_ConnectServiceUnregistered(QtObjectPtr ptr);
void QDBusConnectionInterface_DisconnectServiceUnregistered(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif