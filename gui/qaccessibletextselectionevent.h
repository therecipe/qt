#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent2(QtObjectPtr iface, int start, int end);
QtObjectPtr QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent(QtObjectPtr object, int start, int end);
int QAccessibleTextSelectionEvent_SelectionEnd(QtObjectPtr ptr);
int QAccessibleTextSelectionEvent_SelectionStart(QtObjectPtr ptr);
void QAccessibleTextSelectionEvent_SetSelection(QtObjectPtr ptr, int start, int end);

#ifdef __cplusplus
}
#endif