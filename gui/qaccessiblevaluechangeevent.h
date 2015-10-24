#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent2(QtObjectPtr iface, char* val);
QtObjectPtr QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent(QtObjectPtr object, char* value);
void QAccessibleValueChangeEvent_SetValue(QtObjectPtr ptr, char* value);
char* QAccessibleValueChangeEvent_Value(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif