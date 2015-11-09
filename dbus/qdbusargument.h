#ifdef __cplusplus
extern "C" {
#endif

void* QDBusArgument_NewQDBusArgument();
void* QDBusArgument_NewQDBusArgument2(void* other);
void* QDBusArgument_AsVariant(void* ptr);
int QDBusArgument_AtEnd(void* ptr);
void QDBusArgument_BeginArray(void* ptr, int id);
void QDBusArgument_BeginArray2(void* ptr);
void QDBusArgument_BeginMap(void* ptr, int kid, int vid);
void QDBusArgument_BeginMap2(void* ptr);
void QDBusArgument_BeginMapEntry(void* ptr);
void QDBusArgument_BeginMapEntry2(void* ptr);
void QDBusArgument_BeginStructure(void* ptr);
void QDBusArgument_BeginStructure2(void* ptr);
int QDBusArgument_CurrentType(void* ptr);
void QDBusArgument_EndArray(void* ptr);
void QDBusArgument_EndArray2(void* ptr);
void QDBusArgument_EndMap(void* ptr);
void QDBusArgument_EndMap2(void* ptr);
void QDBusArgument_EndMapEntry(void* ptr);
void QDBusArgument_EndMapEntry2(void* ptr);
void QDBusArgument_EndStructure(void* ptr);
void QDBusArgument_EndStructure2(void* ptr);
void QDBusArgument_DestroyQDBusArgument(void* ptr);

#ifdef __cplusplus
}
#endif