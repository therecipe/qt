#ifdef __cplusplus
extern "C" {
#endif

void* QScriptEngineAgent_NewQScriptEngineAgent(void* engine);
void QScriptEngineAgent_ContextPop(void* ptr);
void QScriptEngineAgent_ContextPush(void* ptr);
void* QScriptEngineAgent_Engine(void* ptr);
void* QScriptEngineAgent_Extension(void* ptr, int extension, void* argument);
int QScriptEngineAgent_SupportsExtension(void* ptr, int extension);
void QScriptEngineAgent_DestroyQScriptEngineAgent(void* ptr);

#ifdef __cplusplus
}
#endif