#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QAccessibleActionInterface_LocalizedActionDescription(QtObjectPtr ptr, char* actionName);
char* QAccessibleActionInterface_LocalizedActionName(QtObjectPtr ptr, char* actionName);
char* QAccessibleActionInterface_ActionNames(QtObjectPtr ptr);
char* QAccessibleActionInterface_QAccessibleActionInterface_DecreaseAction();
void QAccessibleActionInterface_DoAction(QtObjectPtr ptr, char* actionName);
char* QAccessibleActionInterface_QAccessibleActionInterface_IncreaseAction();
char* QAccessibleActionInterface_KeyBindingsForAction(QtObjectPtr ptr, char* actionName);
char* QAccessibleActionInterface_QAccessibleActionInterface_NextPageAction();
char* QAccessibleActionInterface_QAccessibleActionInterface_PressAction();
char* QAccessibleActionInterface_QAccessibleActionInterface_PreviousPageAction();
char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollDownAction();
char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollLeftAction();
char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollRightAction();
char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollUpAction();
char* QAccessibleActionInterface_QAccessibleActionInterface_SetFocusAction();
char* QAccessibleActionInterface_QAccessibleActionInterface_ShowMenuAction();
char* QAccessibleActionInterface_QAccessibleActionInterface_ToggleAction();
void QAccessibleActionInterface_DestroyQAccessibleActionInterface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif