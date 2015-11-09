#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleEvent_NewQAccessibleEvent2(void* interfa, int ty);
void* QAccessibleEvent_NewQAccessibleEvent(void* object, int ty);
void* QAccessibleEvent_AccessibleInterface(void* ptr);
int QAccessibleEvent_Child(void* ptr);
void* QAccessibleEvent_Object(void* ptr);
void QAccessibleEvent_SetChild(void* ptr, int child);
int QAccessibleEvent_Type(void* ptr);
void QAccessibleEvent_DestroyQAccessibleEvent(void* ptr);

#ifdef __cplusplus
}
#endif