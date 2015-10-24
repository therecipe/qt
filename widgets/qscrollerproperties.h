#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScrollerProperties_NewQScrollerProperties();
QtObjectPtr QScrollerProperties_NewQScrollerProperties2(QtObjectPtr sp);
char* QScrollerProperties_ScrollMetric(QtObjectPtr ptr, int metric);
void QScrollerProperties_QScrollerProperties_SetDefaultScrollerProperties(QtObjectPtr sp);
void QScrollerProperties_SetScrollMetric(QtObjectPtr ptr, int metric, char* value);
void QScrollerProperties_QScrollerProperties_UnsetDefaultScrollerProperties();
void QScrollerProperties_DestroyQScrollerProperties(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif