#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent2(void* iface, int position, char* text);
void* QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent(void* object, int position, char* text);
int QAccessibleTextRemoveEvent_ChangePosition(void* ptr);
char* QAccessibleTextRemoveEvent_TextRemoved(void* ptr);

#ifdef __cplusplus
}
#endif