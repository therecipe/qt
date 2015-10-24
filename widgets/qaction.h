#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAction_AutoRepeat(QtObjectPtr ptr);
char* QAction_IconText(QtObjectPtr ptr);
int QAction_IsCheckable(QtObjectPtr ptr);
int QAction_IsChecked(QtObjectPtr ptr);
int QAction_IsEnabled(QtObjectPtr ptr);
int QAction_IsIconVisibleInMenu(QtObjectPtr ptr);
int QAction_IsVisible(QtObjectPtr ptr);
int QAction_MenuRole(QtObjectPtr ptr);
int QAction_Priority(QtObjectPtr ptr);
void QAction_SetAutoRepeat(QtObjectPtr ptr, int v);
void QAction_SetCheckable(QtObjectPtr ptr, int v);
void QAction_SetChecked(QtObjectPtr ptr, int v);
void QAction_SetData(QtObjectPtr ptr, char* userData);
void QAction_SetEnabled(QtObjectPtr ptr, int v);
void QAction_SetFont(QtObjectPtr ptr, QtObjectPtr font);
void QAction_SetIcon(QtObjectPtr ptr, QtObjectPtr icon);
void QAction_SetIconText(QtObjectPtr ptr, char* text);
void QAction_SetIconVisibleInMenu(QtObjectPtr ptr, int visible);
void QAction_SetMenuRole(QtObjectPtr ptr, int menuRole);
void QAction_SetPriority(QtObjectPtr ptr, int priority);
void QAction_SetShortcut(QtObjectPtr ptr, QtObjectPtr shortcut);
void QAction_SetShortcutContext(QtObjectPtr ptr, int context);
void QAction_SetStatusTip(QtObjectPtr ptr, char* statusTip);
void QAction_SetText(QtObjectPtr ptr, char* text);
void QAction_SetToolTip(QtObjectPtr ptr, char* tip);
void QAction_SetVisible(QtObjectPtr ptr, int v);
void QAction_SetWhatsThis(QtObjectPtr ptr, char* what);
int QAction_ShortcutContext(QtObjectPtr ptr);
char* QAction_StatusTip(QtObjectPtr ptr);
char* QAction_Text(QtObjectPtr ptr);
void QAction_Toggle(QtObjectPtr ptr);
char* QAction_ToolTip(QtObjectPtr ptr);
char* QAction_WhatsThis(QtObjectPtr ptr);
QtObjectPtr QAction_NewQAction(QtObjectPtr parent);
QtObjectPtr QAction_NewQAction3(QtObjectPtr icon, char* text, QtObjectPtr parent);
QtObjectPtr QAction_NewQAction2(char* text, QtObjectPtr parent);
QtObjectPtr QAction_ActionGroup(QtObjectPtr ptr);
void QAction_Activate(QtObjectPtr ptr, int event);
void QAction_ConnectChanged(QtObjectPtr ptr);
void QAction_DisconnectChanged(QtObjectPtr ptr);
char* QAction_Data(QtObjectPtr ptr);
void QAction_Hover(QtObjectPtr ptr);
void QAction_ConnectHovered(QtObjectPtr ptr);
void QAction_DisconnectHovered(QtObjectPtr ptr);
int QAction_IsSeparator(QtObjectPtr ptr);
QtObjectPtr QAction_Menu(QtObjectPtr ptr);
QtObjectPtr QAction_ParentWidget(QtObjectPtr ptr);
void QAction_SetActionGroup(QtObjectPtr ptr, QtObjectPtr group);
void QAction_SetDisabled(QtObjectPtr ptr, int b);
void QAction_SetMenu(QtObjectPtr ptr, QtObjectPtr menu);
void QAction_SetSeparator(QtObjectPtr ptr, int b);
void QAction_SetShortcuts2(QtObjectPtr ptr, int key);
int QAction_ShowStatusText(QtObjectPtr ptr, QtObjectPtr widget);
void QAction_ConnectToggled(QtObjectPtr ptr);
void QAction_DisconnectToggled(QtObjectPtr ptr);
void QAction_Trigger(QtObjectPtr ptr);
void QAction_ConnectTriggered(QtObjectPtr ptr);
void QAction_DisconnectTriggered(QtObjectPtr ptr);
void QAction_DestroyQAction(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif