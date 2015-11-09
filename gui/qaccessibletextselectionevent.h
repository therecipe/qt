#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent2(void* iface, int start, int end);
void* QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent(void* object, int start, int end);
int QAccessibleTextSelectionEvent_SelectionEnd(void* ptr);
int QAccessibleTextSelectionEvent_SelectionStart(void* ptr);
void QAccessibleTextSelectionEvent_SetSelection(void* ptr, int start, int end);

#ifdef __cplusplus
}
#endif