#ifdef __cplusplus
extern "C" {
#endif

void* QScriptClass_NewQScriptClass(void* engine);
void* QScriptClass_Engine(void* ptr);
void* QScriptClass_Extension(void* ptr, int extension, void* argument);
char* QScriptClass_Name(void* ptr);
void* QScriptClass_NewIterator(void* ptr, void* object);
void* QScriptClass_Prototype(void* ptr);
int QScriptClass_SupportsExtension(void* ptr, int extension);
void QScriptClass_DestroyQScriptClass(void* ptr);

#ifdef __cplusplus
}
#endif