#ifdef __cplusplus
extern "C" {
#endif

void* QJSEngine_NewQJSEngine();
void* QJSEngine_NewQJSEngine2(void* parent);
void QJSEngine_CollectGarbage(void* ptr);
void* QJSEngine_Evaluate(void* ptr, char* program, char* fileName, int lineNumber);
void* QJSEngine_GlobalObject(void* ptr);
void QJSEngine_InstallTranslatorFunctions(void* ptr, void* object);
void* QJSEngine_NewObject(void* ptr);
void* QJSEngine_NewQObject(void* ptr, void* object);
void QJSEngine_DestroyQJSEngine(void* ptr);

#ifdef __cplusplus
}
#endif