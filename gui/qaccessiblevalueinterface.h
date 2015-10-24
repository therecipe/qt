#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QAccessibleValueInterface_CurrentValue(QtObjectPtr ptr);
char* QAccessibleValueInterface_MaximumValue(QtObjectPtr ptr);
char* QAccessibleValueInterface_MinimumStepSize(QtObjectPtr ptr);
char* QAccessibleValueInterface_MinimumValue(QtObjectPtr ptr);
void QAccessibleValueInterface_SetCurrentValue(QtObjectPtr ptr, char* value);
void QAccessibleValueInterface_DestroyQAccessibleValueInterface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif