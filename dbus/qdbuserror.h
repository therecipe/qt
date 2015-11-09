#ifdef __cplusplus
extern "C" {
#endif

char* QDBusError_QDBusError_ErrorString(int error);
int QDBusError_IsValid(void* ptr);
char* QDBusError_Message(void* ptr);
char* QDBusError_Name(void* ptr);
int QDBusError_Type(void* ptr);

#ifdef __cplusplus
}
#endif