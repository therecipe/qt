#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlScriptString_NewQQmlScriptString();
QtObjectPtr QQmlScriptString_NewQQmlScriptString2(QtObjectPtr other);
int QQmlScriptString_BooleanLiteral(QtObjectPtr ptr, int ok);
int QQmlScriptString_IsEmpty(QtObjectPtr ptr);
int QQmlScriptString_IsNullLiteral(QtObjectPtr ptr);
int QQmlScriptString_IsUndefinedLiteral(QtObjectPtr ptr);
char* QQmlScriptString_StringLiteral(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif