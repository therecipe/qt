#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent2(QtObjectPtr iface, int position, char* text);
QtObjectPtr QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent(QtObjectPtr object, int position, char* text);
int QAccessibleTextRemoveEvent_ChangePosition(QtObjectPtr ptr);
char* QAccessibleTextRemoveEvent_TextRemoved(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif