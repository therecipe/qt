#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent2(void* iface, int position, char* oldText, char* text);
void* QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent(void* object, int position, char* oldText, char* text);
int QAccessibleTextUpdateEvent_ChangePosition(void* ptr);
char* QAccessibleTextUpdateEvent_TextInserted(void* ptr);
char* QAccessibleTextUpdateEvent_TextRemoved(void* ptr);

#ifdef __cplusplus
}
#endif