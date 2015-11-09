#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleWidget_NewQAccessibleWidget(void* w, int role, char* name);
char* QAccessibleWidget_ActionNames(void* ptr);
void* QAccessibleWidget_BackgroundColor(void* ptr);
void* QAccessibleWidget_Child(void* ptr, int index);
int QAccessibleWidget_ChildCount(void* ptr);
void QAccessibleWidget_DoAction(void* ptr, char* actionName);
void* QAccessibleWidget_FocusChild(void* ptr);
void* QAccessibleWidget_ForegroundColor(void* ptr);
int QAccessibleWidget_IndexOfChild(void* ptr, void* child);
void* QAccessibleWidget_Interface_cast(void* ptr, int t);
int QAccessibleWidget_IsValid(void* ptr);
char* QAccessibleWidget_KeyBindingsForAction(void* ptr, char* actionName);
void* QAccessibleWidget_Parent(void* ptr);
int QAccessibleWidget_Role(void* ptr);
char* QAccessibleWidget_Text(void* ptr, int t);
void* QAccessibleWidget_Window(void* ptr);

#ifdef __cplusplus
}
#endif