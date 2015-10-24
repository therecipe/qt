#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent2(QtObjectPtr iface, int cursorPos);
QtObjectPtr QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent(QtObjectPtr object, int cursorPos);
int QAccessibleTextCursorEvent_CursorPosition(QtObjectPtr ptr);
void QAccessibleTextCursorEvent_SetCursorPosition(QtObjectPtr ptr, int position);

#ifdef __cplusplus
}
#endif