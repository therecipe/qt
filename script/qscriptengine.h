#ifdef __cplusplus
extern "C" {
#endif

void* QScriptEngine_NewQScriptEngine();
void* QScriptEngine_NewQScriptEngine2(void* parent);
void QScriptEngine_AbortEvaluation(void* ptr, void* result);
void* QScriptEngine_Agent(void* ptr);
char* QScriptEngine_AvailableExtensions(void* ptr);
void QScriptEngine_ClearExceptions(void* ptr);
void QScriptEngine_CollectGarbage(void* ptr);
void* QScriptEngine_CurrentContext(void* ptr);
void* QScriptEngine_DefaultPrototype(void* ptr, int metaTypeId);
void* QScriptEngine_Evaluate2(void* ptr, void* program);
void* QScriptEngine_Evaluate(void* ptr, char* program, char* fileName, int lineNumber);
void* QScriptEngine_GlobalObject(void* ptr);
int QScriptEngine_HasUncaughtException(void* ptr);
void* QScriptEngine_ImportExtension(void* ptr, char* extension);
char* QScriptEngine_ImportedExtensions(void* ptr);
void QScriptEngine_InstallTranslatorFunctions(void* ptr, void* object);
int QScriptEngine_IsEvaluating(void* ptr);
void* QScriptEngine_NewDate2(void* ptr, void* value);
void* QScriptEngine_NewObject(void* ptr);
void* QScriptEngine_NewObject2(void* ptr, void* scriptClass, void* data);
void* QScriptEngine_NewQMetaObject(void* ptr, void* metaObject, void* ctor);
void* QScriptEngine_NewQObject(void* ptr, void* object, int ownership, int options);
void* QScriptEngine_NewQObject2(void* ptr, void* scriptObject, void* qtObject, int ownership, int options);
void* QScriptEngine_NewRegExp(void* ptr, void* regexp);
void* QScriptEngine_NewRegExp2(void* ptr, char* pattern, char* flags);
void* QScriptEngine_NewVariant2(void* ptr, void* object, void* value);
void* QScriptEngine_NewVariant(void* ptr, void* value);
void* QScriptEngine_NullValue(void* ptr);
void QScriptEngine_PopContext(void* ptr);
int QScriptEngine_ProcessEventsInterval(void* ptr);
void* QScriptEngine_PushContext(void* ptr);
void QScriptEngine_ReportAdditionalMemoryCost(void* ptr, int size);
void QScriptEngine_SetAgent(void* ptr, void* agent);
void QScriptEngine_SetDefaultPrototype(void* ptr, int metaTypeId, void* prototype);
void QScriptEngine_SetGlobalObject(void* ptr, void* object);
void QScriptEngine_SetProcessEventsInterval(void* ptr, int interval);
void QScriptEngine_ConnectSignalHandlerException(void* ptr);
void QScriptEngine_DisconnectSignalHandlerException(void* ptr);
void* QScriptEngine_ToObject(void* ptr, void* value);
void* QScriptEngine_UncaughtException(void* ptr);
char* QScriptEngine_UncaughtExceptionBacktrace(void* ptr);
int QScriptEngine_UncaughtExceptionLineNumber(void* ptr);
void* QScriptEngine_UndefinedValue(void* ptr);
void QScriptEngine_DestroyQScriptEngine(void* ptr);

#ifdef __cplusplus
}
#endif