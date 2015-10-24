#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QDBusError_QDBusError_ErrorString(int error);
int QDBusError_IsValid(QtObjectPtr ptr);
char* QDBusError_Message(QtObjectPtr ptr);
char* QDBusError_Name(QtObjectPtr ptr);
int QDBusError_Type(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif