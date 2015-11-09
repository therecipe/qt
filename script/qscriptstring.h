#ifdef __cplusplus
extern "C" {
#endif

void* QScriptString_NewQScriptString();
void* QScriptString_NewQScriptString2(void* other);
int QScriptString_IsValid(void* ptr);
char* QScriptString_ToString(void* ptr);
void QScriptString_DestroyQScriptString(void* ptr);

#ifdef __cplusplus
}
#endif