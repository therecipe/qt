#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleObject_ChildAt(void* ptr, int x, int y);
int QAccessibleObject_IsValid(void* ptr);
void* QAccessibleObject_Object(void* ptr);
void QAccessibleObject_SetText(void* ptr, int t, char* text);

#ifdef __cplusplus
}
#endif