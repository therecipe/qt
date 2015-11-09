#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent2(void* iface, int cursorPos);
void* QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent(void* object, int cursorPos);
int QAccessibleTextCursorEvent_CursorPosition(void* ptr);
void QAccessibleTextCursorEvent_SetCursorPosition(void* ptr, int position);

#ifdef __cplusplus
}
#endif