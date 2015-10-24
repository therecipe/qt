#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleInterface_ActionInterface(QtObjectPtr ptr);
QtObjectPtr QAccessibleInterface_Child(QtObjectPtr ptr, int index);
QtObjectPtr QAccessibleInterface_ChildAt(QtObjectPtr ptr, int x, int y);
int QAccessibleInterface_ChildCount(QtObjectPtr ptr);
QtObjectPtr QAccessibleInterface_FocusChild(QtObjectPtr ptr);
int QAccessibleInterface_IndexOfChild(QtObjectPtr ptr, QtObjectPtr child);
void QAccessibleInterface_Interface_cast(QtObjectPtr ptr, int ty);
int QAccessibleInterface_IsValid(QtObjectPtr ptr);
QtObjectPtr QAccessibleInterface_Object(QtObjectPtr ptr);
QtObjectPtr QAccessibleInterface_Parent(QtObjectPtr ptr);
int QAccessibleInterface_Role(QtObjectPtr ptr);
void QAccessibleInterface_SetText(QtObjectPtr ptr, int t, char* text);
QtObjectPtr QAccessibleInterface_TableCellInterface(QtObjectPtr ptr);
QtObjectPtr QAccessibleInterface_TableInterface(QtObjectPtr ptr);
char* QAccessibleInterface_Text(QtObjectPtr ptr, int t);
QtObjectPtr QAccessibleInterface_TextInterface(QtObjectPtr ptr);
QtObjectPtr QAccessibleInterface_ValueInterface(QtObjectPtr ptr);
QtObjectPtr QAccessibleInterface_Window(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif