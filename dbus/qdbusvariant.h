#ifdef __cplusplus
extern "C" {
#endif

void* QDBusVariant_NewQDBusVariant();
void* QDBusVariant_NewQDBusVariant2(void* variant);
void QDBusVariant_SetVariant(void* ptr, void* variant);
void* QDBusVariant_Variant(void* ptr);

#ifdef __cplusplus
}
#endif