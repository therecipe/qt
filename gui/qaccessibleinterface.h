#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleInterface_ActionInterface(void* ptr);
void* QAccessibleInterface_BackgroundColor(void* ptr);
void* QAccessibleInterface_Child(void* ptr, int index);
void* QAccessibleInterface_ChildAt(void* ptr, int x, int y);
int QAccessibleInterface_ChildCount(void* ptr);
void* QAccessibleInterface_FocusChild(void* ptr);
void* QAccessibleInterface_ForegroundColor(void* ptr);
int QAccessibleInterface_IndexOfChild(void* ptr, void* child);
void* QAccessibleInterface_Interface_cast(void* ptr, int ty);
int QAccessibleInterface_IsValid(void* ptr);
void* QAccessibleInterface_Object(void* ptr);
void* QAccessibleInterface_Parent(void* ptr);
int QAccessibleInterface_Role(void* ptr);
void QAccessibleInterface_SetText(void* ptr, int t, char* text);
void* QAccessibleInterface_TableCellInterface(void* ptr);
void* QAccessibleInterface_TableInterface(void* ptr);
char* QAccessibleInterface_Text(void* ptr, int t);
void* QAccessibleInterface_TextInterface(void* ptr);
void* QAccessibleInterface_ValueInterface(void* ptr);
void* QAccessibleInterface_Window(void* ptr);

#ifdef __cplusplus
}
#endif