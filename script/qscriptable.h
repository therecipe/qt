#ifdef __cplusplus
extern "C" {
#endif

void* QScriptable_Argument(void* ptr, int index);
int QScriptable_ArgumentCount(void* ptr);
void* QScriptable_Context(void* ptr);
void* QScriptable_Engine(void* ptr);
void* QScriptable_ThisObject(void* ptr);

#ifdef __cplusplus
}
#endif