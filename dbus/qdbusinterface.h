#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusInterface_NewQDBusInterface(char* service, char* path, char* interfa, QtObjectPtr connection, QtObjectPtr parent);
void QDBusInterface_DestroyQDBusInterface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif