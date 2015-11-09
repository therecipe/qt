#ifdef __cplusplus
extern "C" {
#endif

void* QScriptContext_ActivationObject(void* ptr);
void* QScriptContext_Argument(void* ptr, int index);
int QScriptContext_ArgumentCount(void* ptr);
void* QScriptContext_ArgumentsObject(void* ptr);
char* QScriptContext_Backtrace(void* ptr);
void* QScriptContext_Callee(void* ptr);
void* QScriptContext_Engine(void* ptr);
int QScriptContext_IsCalledAsConstructor(void* ptr);
void* QScriptContext_ParentContext(void* ptr);
void QScriptContext_SetActivationObject(void* ptr, void* activation);
void QScriptContext_SetThisObject(void* ptr, void* thisObject);
void* QScriptContext_ThisObject(void* ptr);
void* QScriptContext_ThrowError(void* ptr, int error, char* text);
void* QScriptContext_ThrowError2(void* ptr, char* text);
void* QScriptContext_ThrowValue(void* ptr, void* value);
char* QScriptContext_ToString(void* ptr);
void QScriptContext_DestroyQScriptContext(void* ptr);

#ifdef __cplusplus
}
#endif