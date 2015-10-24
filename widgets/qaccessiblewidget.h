#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleWidget_NewQAccessibleWidget(QtObjectPtr w, int role, char* name);
char* QAccessibleWidget_ActionNames(QtObjectPtr ptr);
QtObjectPtr QAccessibleWidget_Child(QtObjectPtr ptr, int index);
int QAccessibleWidget_ChildCount(QtObjectPtr ptr);
void QAccessibleWidget_DoAction(QtObjectPtr ptr, char* actionName);
QtObjectPtr QAccessibleWidget_FocusChild(QtObjectPtr ptr);
int QAccessibleWidget_IndexOfChild(QtObjectPtr ptr, QtObjectPtr child);
void QAccessibleWidget_Interface_cast(QtObjectPtr ptr, int t);
int QAccessibleWidget_IsValid(QtObjectPtr ptr);
char* QAccessibleWidget_KeyBindingsForAction(QtObjectPtr ptr, char* actionName);
QtObjectPtr QAccessibleWidget_Parent(QtObjectPtr ptr);
int QAccessibleWidget_Role(QtObjectPtr ptr);
char* QAccessibleWidget_Text(QtObjectPtr ptr, int t);
QtObjectPtr QAccessibleWidget_Window(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif