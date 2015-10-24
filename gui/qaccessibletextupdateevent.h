#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent2(QtObjectPtr iface, int position, char* oldText, char* text);
QtObjectPtr QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent(QtObjectPtr object, int position, char* oldText, char* text);
int QAccessibleTextUpdateEvent_ChangePosition(QtObjectPtr ptr);
char* QAccessibleTextUpdateEvent_TextInserted(QtObjectPtr ptr);
char* QAccessibleTextUpdateEvent_TextRemoved(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif