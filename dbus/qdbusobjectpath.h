#ifdef __cplusplus
extern "C" {
#endif

void* QDBusObjectPath_NewQDBusObjectPath();
void* QDBusObjectPath_NewQDBusObjectPath3(void* path);
void* QDBusObjectPath_NewQDBusObjectPath4(char* path);
void* QDBusObjectPath_NewQDBusObjectPath2(char* path);
char* QDBusObjectPath_Path(void* ptr);
void QDBusObjectPath_SetPath(void* ptr, char* path);

#ifdef __cplusplus
}
#endif