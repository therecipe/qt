#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QMenu_New_QWidget(QtObjectPtr parent);
QtObjectPtr QMenu_New_String_QWidget(char* title, QtObjectPtr parent);
void QMenu_Destroy(QtObjectPtr ptr);
QtObjectPtr QMenu_ActiveAction(QtObjectPtr ptr);
QtObjectPtr QMenu_AddAction_String(QtObjectPtr ptr, char* text);
void QMenu_AddAction_QAction(QtObjectPtr ptr, QtObjectPtr action);
QtObjectPtr QMenu_AddMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu);
QtObjectPtr QMenu_AddMenu_String(QtObjectPtr ptr, char* title);
QtObjectPtr QMenu_AddSection_String(QtObjectPtr ptr, char* text);
QtObjectPtr QMenu_AddSeparator(QtObjectPtr ptr);
void QMenu_Clear(QtObjectPtr ptr);
QtObjectPtr QMenu_DefaultAction(QtObjectPtr ptr);
QtObjectPtr QMenu_Exec(QtObjectPtr ptr);
void QMenu_HideTearOffMenu(QtObjectPtr ptr);
QtObjectPtr QMenu_InsertMenu_QAction_QMenu(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr menu);
QtObjectPtr QMenu_InsertSection_QAction_String(QtObjectPtr ptr, QtObjectPtr before, char* text);
QtObjectPtr QMenu_InsertSeparator_QAction(QtObjectPtr ptr, QtObjectPtr before);
int QMenu_IsEmpty(QtObjectPtr ptr);
int QMenu_IsTearOffEnabled(QtObjectPtr ptr);
int QMenu_IsTearOffMenuVisible(QtObjectPtr ptr);
QtObjectPtr QMenu_MenuAction(QtObjectPtr ptr);
int QMenu_SeparatorsCollapsible(QtObjectPtr ptr);
void QMenu_SetActiveAction_QAction(QtObjectPtr ptr, QtObjectPtr act);
void QMenu_SetDefaultAction_QAction(QtObjectPtr ptr, QtObjectPtr act);
void QMenu_SetSeparatorsCollapsible_Bool(QtObjectPtr ptr, int collapse);
void QMenu_SetTearOffEnabled_Bool(QtObjectPtr ptr, int tearOffEnabled);
void QMenu_SetTitle_String(QtObjectPtr ptr, char* title);
void QMenu_SetToolTipsVisible_Bool(QtObjectPtr ptr, int visible);
char* QMenu_Title(QtObjectPtr ptr);
int QMenu_ToolTipsVisible(QtObjectPtr ptr);
//Signals
void QMenu_ConnectSignalAboutToHide(QtObjectPtr ptr);
void QMenu_DisconnectSignalAboutToHide(QtObjectPtr ptr);
void QMenu_ConnectSignalAboutToShow(QtObjectPtr ptr);
void QMenu_DisconnectSignalAboutToShow(QtObjectPtr ptr);
void QMenu_ConnectSignalHovered(QtObjectPtr ptr);
void QMenu_DisconnectSignalHovered(QtObjectPtr ptr);
void QMenu_ConnectSignalTriggered(QtObjectPtr ptr);
void QMenu_DisconnectSignalTriggered(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
