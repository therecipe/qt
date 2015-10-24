#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMenu_IsTearOffEnabled(QtObjectPtr ptr);
int QMenu_SeparatorsCollapsible(QtObjectPtr ptr);
void QMenu_SetIcon(QtObjectPtr ptr, QtObjectPtr icon);
void QMenu_SetSeparatorsCollapsible(QtObjectPtr ptr, int collapse);
void QMenu_SetTearOffEnabled(QtObjectPtr ptr, int v);
void QMenu_SetTitle(QtObjectPtr ptr, char* title);
void QMenu_SetToolTipsVisible(QtObjectPtr ptr, int visible);
char* QMenu_Title(QtObjectPtr ptr);
int QMenu_ToolTipsVisible(QtObjectPtr ptr);
QtObjectPtr QMenu_NewQMenu(QtObjectPtr parent);
QtObjectPtr QMenu_NewQMenu2(char* title, QtObjectPtr parent);
void QMenu_ConnectAboutToHide(QtObjectPtr ptr);
void QMenu_DisconnectAboutToHide(QtObjectPtr ptr);
void QMenu_ConnectAboutToShow(QtObjectPtr ptr);
void QMenu_DisconnectAboutToShow(QtObjectPtr ptr);
QtObjectPtr QMenu_ActionAt(QtObjectPtr ptr, QtObjectPtr pt);
QtObjectPtr QMenu_ActiveAction(QtObjectPtr ptr);
QtObjectPtr QMenu_AddAction2(QtObjectPtr ptr, QtObjectPtr icon, char* text);
QtObjectPtr QMenu_AddAction4(QtObjectPtr ptr, QtObjectPtr icon, char* text, QtObjectPtr receiver, char* member, QtObjectPtr shortcut);
QtObjectPtr QMenu_AddAction(QtObjectPtr ptr, char* text);
QtObjectPtr QMenu_AddAction3(QtObjectPtr ptr, char* text, QtObjectPtr receiver, char* member, QtObjectPtr shortcut);
QtObjectPtr QMenu_AddMenu(QtObjectPtr ptr, QtObjectPtr menu);
QtObjectPtr QMenu_AddMenu3(QtObjectPtr ptr, QtObjectPtr icon, char* title);
QtObjectPtr QMenu_AddMenu2(QtObjectPtr ptr, char* title);
QtObjectPtr QMenu_AddSection2(QtObjectPtr ptr, QtObjectPtr icon, char* text);
QtObjectPtr QMenu_AddSection(QtObjectPtr ptr, char* text);
QtObjectPtr QMenu_AddSeparator(QtObjectPtr ptr);
void QMenu_Clear(QtObjectPtr ptr);
QtObjectPtr QMenu_Exec(QtObjectPtr ptr);
QtObjectPtr QMenu_Exec2(QtObjectPtr ptr, QtObjectPtr p, QtObjectPtr action);
void QMenu_HideTearOffMenu(QtObjectPtr ptr);
void QMenu_ConnectHovered(QtObjectPtr ptr);
void QMenu_DisconnectHovered(QtObjectPtr ptr);
QtObjectPtr QMenu_InsertMenu(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr menu);
QtObjectPtr QMenu_InsertSection2(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr icon, char* text);
QtObjectPtr QMenu_InsertSection(QtObjectPtr ptr, QtObjectPtr before, char* text);
QtObjectPtr QMenu_InsertSeparator(QtObjectPtr ptr, QtObjectPtr before);
int QMenu_IsEmpty(QtObjectPtr ptr);
int QMenu_IsTearOffMenuVisible(QtObjectPtr ptr);
QtObjectPtr QMenu_MenuAction(QtObjectPtr ptr);
void QMenu_Popup(QtObjectPtr ptr, QtObjectPtr p, QtObjectPtr atAction);
void QMenu_SetActiveAction(QtObjectPtr ptr, QtObjectPtr act);
void QMenu_ConnectTriggered(QtObjectPtr ptr);
void QMenu_DisconnectTriggered(QtObjectPtr ptr);
void QMenu_DestroyQMenu(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif