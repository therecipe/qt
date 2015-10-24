#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusObjectPath_NewQDBusObjectPath();
QtObjectPtr QDBusObjectPath_NewQDBusObjectPath3(QtObjectPtr path);
QtObjectPtr QDBusObjectPath_NewQDBusObjectPath4(char* path);
QtObjectPtr QDBusObjectPath_NewQDBusObjectPath2(char* path);
char* QDBusObjectPath_Path(QtObjectPtr ptr);
void QDBusObjectPath_SetPath(QtObjectPtr ptr, char* path);

#ifdef __cplusplus
}
#endif