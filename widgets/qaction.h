#ifdef __cplusplus
extern "C" {
#endif

int QAction_AutoRepeat(void* ptr);
char* QAction_IconText(void* ptr);
int QAction_IsCheckable(void* ptr);
int QAction_IsChecked(void* ptr);
int QAction_IsEnabled(void* ptr);
int QAction_IsIconVisibleInMenu(void* ptr);
int QAction_IsVisible(void* ptr);
int QAction_MenuRole(void* ptr);
int QAction_Priority(void* ptr);
void QAction_SetAutoRepeat(void* ptr, int v);
void QAction_SetCheckable(void* ptr, int v);
void QAction_SetChecked(void* ptr, int v);
void QAction_SetData(void* ptr, void* userData);
void QAction_SetEnabled(void* ptr, int v);
void QAction_SetFont(void* ptr, void* font);
void QAction_SetIcon(void* ptr, void* icon);
void QAction_SetIconText(void* ptr, char* text);
void QAction_SetIconVisibleInMenu(void* ptr, int visible);
void QAction_SetMenuRole(void* ptr, int menuRole);
void QAction_SetPriority(void* ptr, int priority);
void QAction_SetShortcut(void* ptr, void* shortcut);
void QAction_SetShortcutContext(void* ptr, int context);
void QAction_SetStatusTip(void* ptr, char* statusTip);
void QAction_SetText(void* ptr, char* text);
void QAction_SetToolTip(void* ptr, char* tip);
void QAction_SetVisible(void* ptr, int v);
void QAction_SetWhatsThis(void* ptr, char* what);
int QAction_ShortcutContext(void* ptr);
char* QAction_StatusTip(void* ptr);
char* QAction_Text(void* ptr);
void QAction_Toggle(void* ptr);
char* QAction_ToolTip(void* ptr);
char* QAction_WhatsThis(void* ptr);
void* QAction_NewQAction(void* parent);
void* QAction_NewQAction3(void* icon, char* text, void* parent);
void* QAction_NewQAction2(char* text, void* parent);
void* QAction_ActionGroup(void* ptr);
void QAction_Activate(void* ptr, int event);
void QAction_ConnectChanged(void* ptr);
void QAction_DisconnectChanged(void* ptr);
void* QAction_Data(void* ptr);
void QAction_Hover(void* ptr);
void QAction_ConnectHovered(void* ptr);
void QAction_DisconnectHovered(void* ptr);
int QAction_IsSeparator(void* ptr);
void* QAction_Menu(void* ptr);
void* QAction_ParentWidget(void* ptr);
void QAction_SetActionGroup(void* ptr, void* group);
void QAction_SetDisabled(void* ptr, int b);
void QAction_SetMenu(void* ptr, void* menu);
void QAction_SetSeparator(void* ptr, int b);
void QAction_SetShortcuts2(void* ptr, int key);
int QAction_ShowStatusText(void* ptr, void* widget);
void QAction_ConnectToggled(void* ptr);
void QAction_DisconnectToggled(void* ptr);
void QAction_Trigger(void* ptr);
void QAction_ConnectTriggered(void* ptr);
void QAction_DisconnectTriggered(void* ptr);
void QAction_DestroyQAction(void* ptr);

#ifdef __cplusplus
}
#endif