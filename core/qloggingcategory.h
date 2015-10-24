#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QLoggingCategory_NewQLoggingCategory(char* category);
QtObjectPtr QLoggingCategory_QLoggingCategory_DefaultCategory();
int QLoggingCategory_IsCriticalEnabled(QtObjectPtr ptr);
int QLoggingCategory_IsDebugEnabled(QtObjectPtr ptr);
int QLoggingCategory_IsInfoEnabled(QtObjectPtr ptr);
int QLoggingCategory_IsWarningEnabled(QtObjectPtr ptr);
void QLoggingCategory_QLoggingCategory_SetFilterRules(char* rules);
void QLoggingCategory_DestroyQLoggingCategory(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif