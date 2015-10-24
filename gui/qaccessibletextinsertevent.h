#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent2(QtObjectPtr iface, int position, char* text);
QtObjectPtr QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent(QtObjectPtr object, int position, char* text);
int QAccessibleTextInsertEvent_ChangePosition(QtObjectPtr ptr);
char* QAccessibleTextInsertEvent_TextInserted(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif