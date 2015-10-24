#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor();
QtObjectPtr QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(QtObjectPtr other);
QtObjectPtr QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(int fileDescriptor);
int QDBusUnixFileDescriptor_FileDescriptor(QtObjectPtr ptr);
int QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported();
int QDBusUnixFileDescriptor_IsValid(QtObjectPtr ptr);
void QDBusUnixFileDescriptor_SetFileDescriptor(QtObjectPtr ptr, int fileDescriptor);
void QDBusUnixFileDescriptor_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif