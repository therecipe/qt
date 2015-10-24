#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusVariant_NewQDBusVariant();
QtObjectPtr QDBusVariant_NewQDBusVariant2(char* variant);
void QDBusVariant_SetVariant(QtObjectPtr ptr, char* variant);
char* QDBusVariant_Variant(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif