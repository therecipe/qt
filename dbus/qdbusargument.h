#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusArgument_NewQDBusArgument();
QtObjectPtr QDBusArgument_NewQDBusArgument2(QtObjectPtr other);
char* QDBusArgument_AsVariant(QtObjectPtr ptr);
int QDBusArgument_AtEnd(QtObjectPtr ptr);
void QDBusArgument_BeginArray(QtObjectPtr ptr, int id);
void QDBusArgument_BeginArray2(QtObjectPtr ptr);
void QDBusArgument_BeginMap(QtObjectPtr ptr, int kid, int vid);
void QDBusArgument_BeginMap2(QtObjectPtr ptr);
void QDBusArgument_BeginMapEntry(QtObjectPtr ptr);
void QDBusArgument_BeginMapEntry2(QtObjectPtr ptr);
void QDBusArgument_BeginStructure(QtObjectPtr ptr);
void QDBusArgument_BeginStructure2(QtObjectPtr ptr);
int QDBusArgument_CurrentType(QtObjectPtr ptr);
void QDBusArgument_EndArray(QtObjectPtr ptr);
void QDBusArgument_EndArray2(QtObjectPtr ptr);
void QDBusArgument_EndMap(QtObjectPtr ptr);
void QDBusArgument_EndMap2(QtObjectPtr ptr);
void QDBusArgument_EndMapEntry(QtObjectPtr ptr);
void QDBusArgument_EndMapEntry2(QtObjectPtr ptr);
void QDBusArgument_EndStructure(QtObjectPtr ptr);
void QDBusArgument_EndStructure2(QtObjectPtr ptr);
void QDBusArgument_DestroyQDBusArgument(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif