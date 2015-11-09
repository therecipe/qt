#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent2(void* iface, int position, char* text);
void* QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent(void* object, int position, char* text);
int QAccessibleTextInsertEvent_ChangePosition(void* ptr);
char* QAccessibleTextInsertEvent_TextInserted(void* ptr);

#ifdef __cplusplus
}
#endif