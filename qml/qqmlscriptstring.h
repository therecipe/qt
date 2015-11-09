#ifdef __cplusplus
extern "C" {
#endif

void* QQmlScriptString_NewQQmlScriptString();
void* QQmlScriptString_NewQQmlScriptString2(void* other);
int QQmlScriptString_BooleanLiteral(void* ptr, int ok);
int QQmlScriptString_IsEmpty(void* ptr);
int QQmlScriptString_IsNullLiteral(void* ptr);
int QQmlScriptString_IsUndefinedLiteral(void* ptr);
double QQmlScriptString_NumberLiteral(void* ptr, int ok);
char* QQmlScriptString_StringLiteral(void* ptr);

#ifdef __cplusplus
}
#endif