#ifdef __cplusplus
extern "C" {
#endif

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor();
void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(void* other);
void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(int fileDescriptor);
int QDBusUnixFileDescriptor_FileDescriptor(void* ptr);
int QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported();
int QDBusUnixFileDescriptor_IsValid(void* ptr);
void QDBusUnixFileDescriptor_SetFileDescriptor(void* ptr, int fileDescriptor);
void QDBusUnixFileDescriptor_Swap(void* ptr, void* other);
void QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(void* ptr);

#ifdef __cplusplus
}
#endif