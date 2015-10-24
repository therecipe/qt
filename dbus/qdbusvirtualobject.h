#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDBusVirtualObject_HandleMessage(QtObjectPtr ptr, QtObjectPtr message, QtObjectPtr connection);
char* QDBusVirtualObject_Introspect(QtObjectPtr ptr, char* path);
void QDBusVirtualObject_DestroyQDBusVirtualObject(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif