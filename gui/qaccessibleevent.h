#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleEvent_NewQAccessibleEvent2(QtObjectPtr interfa, int ty);
QtObjectPtr QAccessibleEvent_NewQAccessibleEvent(QtObjectPtr object, int ty);
QtObjectPtr QAccessibleEvent_AccessibleInterface(QtObjectPtr ptr);
int QAccessibleEvent_Child(QtObjectPtr ptr);
QtObjectPtr QAccessibleEvent_Object(QtObjectPtr ptr);
void QAccessibleEvent_SetChild(QtObjectPtr ptr, int child);
int QAccessibleEvent_Type(QtObjectPtr ptr);
void QAccessibleEvent_DestroyQAccessibleEvent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif