#ifdef __cplusplus
extern "C" {
#endif

void* QLoggingCategory_NewQLoggingCategory(char* category);
void* QLoggingCategory_QLoggingCategory_DefaultCategory();
int QLoggingCategory_IsCriticalEnabled(void* ptr);
int QLoggingCategory_IsDebugEnabled(void* ptr);
int QLoggingCategory_IsInfoEnabled(void* ptr);
int QLoggingCategory_IsWarningEnabled(void* ptr);
void QLoggingCategory_QLoggingCategory_SetFilterRules(char* rules);
void QLoggingCategory_DestroyQLoggingCategory(void* ptr);

#ifdef __cplusplus
}
#endif