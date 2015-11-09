#ifdef __cplusplus
extern "C" {
#endif

char* QAccessibleActionInterface_LocalizedActionDescription(void* ptr, char* actionName);
char* QAccessibleActionInterface_LocalizedActionName(void* ptr, char* actionName);
char* QAccessibleActionInterface_ActionNames(void* ptr);
char* QAccessibleActionInterface_QAccessibleActionInterface_DecreaseAction();
void QAccessibleActionInterface_DoAction(void* ptr, char* actionName);
char* QAccessibleActionInterface_QAccessibleActionInterface_IncreaseAction();
char* QAccessibleActionInterface_KeyBindingsForAction(void* ptr, char* actionName);
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
void QAccessibleActionInterface_DestroyQAccessibleActionInterface(void* ptr);

#ifdef __cplusplus
}
#endif