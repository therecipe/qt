#ifdef __cplusplus
extern "C" {
#endif

int QDBusVirtualObject_HandleMessage(void* ptr, void* message, void* connection);
char* QDBusVirtualObject_Introspect(void* ptr, char* path);
void QDBusVirtualObject_DestroyQDBusVirtualObject(void* ptr);

#ifdef __cplusplus
}
#endif