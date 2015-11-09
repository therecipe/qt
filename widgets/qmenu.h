#ifdef __cplusplus
extern "C" {
#endif

int QMenu_IsTearOffEnabled(void* ptr);
int QMenu_SeparatorsCollapsible(void* ptr);
void QMenu_SetIcon(void* ptr, void* icon);
void QMenu_SetSeparatorsCollapsible(void* ptr, int collapse);
void QMenu_SetTearOffEnabled(void* ptr, int v);
void QMenu_SetTitle(void* ptr, char* title);
void QMenu_SetToolTipsVisible(void* ptr, int visible);
char* QMenu_Title(void* ptr);
int QMenu_ToolTipsVisible(void* ptr);
void* QMenu_NewQMenu(void* parent);
void* QMenu_NewQMenu2(char* title, void* parent);
void QMenu_ConnectAboutToHide(void* ptr);
void QMenu_DisconnectAboutToHide(void* ptr);
void QMenu_ConnectAboutToShow(void* ptr);
void QMenu_DisconnectAboutToShow(void* ptr);
void* QMenu_ActionAt(void* ptr, void* pt);
void* QMenu_ActiveAction(void* ptr);
void* QMenu_AddAction2(void* ptr, void* icon, char* text);
void* QMenu_AddAction4(void* ptr, void* icon, char* text, void* receiver, char* member, void* shortcut);
void* QMenu_AddAction(void* ptr, char* text);
void* QMenu_AddAction3(void* ptr, char* text, void* receiver, char* member, void* shortcut);
void* QMenu_AddMenu(void* ptr, void* menu);
void* QMenu_AddMenu3(void* ptr, void* icon, char* title);
void* QMenu_AddMenu2(void* ptr, char* title);
void* QMenu_AddSection2(void* ptr, void* icon, char* text);
void* QMenu_AddSection(void* ptr, char* text);
void* QMenu_AddSeparator(void* ptr);
void QMenu_Clear(void* ptr);
void* QMenu_Exec(void* ptr);
void* QMenu_Exec2(void* ptr, void* p, void* action);
void QMenu_HideTearOffMenu(void* ptr);
void QMenu_ConnectHovered(void* ptr);
void QMenu_DisconnectHovered(void* ptr);
void* QMenu_InsertMenu(void* ptr, void* before, void* menu);
void* QMenu_InsertSection2(void* ptr, void* before, void* icon, char* text);
void* QMenu_InsertSection(void* ptr, void* before, char* text);
void* QMenu_InsertSeparator(void* ptr, void* before);
int QMenu_IsEmpty(void* ptr);
int QMenu_IsTearOffMenuVisible(void* ptr);
void* QMenu_MenuAction(void* ptr);
void QMenu_Popup(void* ptr, void* p, void* atAction);
void QMenu_SetActiveAction(void* ptr, void* act);
void QMenu_ConnectTriggered(void* ptr);
void QMenu_DisconnectTriggered(void* ptr);
void QMenu_DestroyQMenu(void* ptr);

#ifdef __cplusplus
}
#endif