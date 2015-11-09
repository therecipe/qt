#ifdef __cplusplus
extern "C" {
#endif

void* QDBusInterface_NewQDBusInterface(char* service, char* path, char* interfa, void* connection, void* parent);
void QDBusInterface_DestroyQDBusInterface(void* ptr);

#ifdef __cplusplus
}
#endif