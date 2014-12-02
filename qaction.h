#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QAction_New_QObject(QtObjectPtr parent);
QtObjectPtr QAction_New_String_QObject(char* text, QtObjectPtr parent);
void QAction_Destroy(QtObjectPtr ptr);
int QAction_AutoRepeat(QtObjectPtr ptr);
char* QAction_IconText(QtObjectPtr ptr);
int QAction_IsCheckable(QtObjectPtr ptr);
int QAction_IsChecked(QtObjectPtr ptr);
int QAction_IsEnabled(QtObjectPtr ptr);
int QAction_IsIconVisibleInMenu(QtObjectPtr ptr);
int QAction_IsSeparator(QtObjectPtr ptr);
int QAction_IsVisible(QtObjectPtr ptr);
QtObjectPtr QAction_Menu(QtObjectPtr ptr);
QtObjectPtr QAction_ParentWidget(QtObjectPtr ptr);
void QAction_SetAutoRepeat_Bool(QtObjectPtr ptr, int autoRepeat);
void QAction_SetCheckable_Bool(QtObjectPtr ptr, int checkable);
void QAction_SetIconText_String(QtObjectPtr ptr, char* text);
void QAction_SetIconVisibleInMenu_Bool(QtObjectPtr ptr, int visible);
void QAction_SetMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu);
void QAction_SetSeparator_Bool(QtObjectPtr ptr, int bo);
void QAction_SetShortcutContext_ShortcutContext(QtObjectPtr ptr, int context);
void QAction_SetStatusTip_String(QtObjectPtr ptr, char* statusTip);
void QAction_SetText_String(QtObjectPtr ptr, char* text);
void QAction_SetToolTip_String(QtObjectPtr ptr, char* tip);
void QAction_SetWhatsThis_String(QtObjectPtr ptr, char* what);
int QAction_ShortcutContext(QtObjectPtr ptr);
int QAction_ShowStatusText_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
char* QAction_StatusTip(QtObjectPtr ptr);
char* QAction_Text(QtObjectPtr ptr);
char* QAction_ToolTip(QtObjectPtr ptr);
char* QAction_WhatsThis(QtObjectPtr ptr);
//Public Slots
void QAction_ConnectSlotHover(QtObjectPtr ptr);
void QAction_DisconnectSlotHover(QtObjectPtr ptr);
void QAction_Hover(QtObjectPtr ptr);
void QAction_ConnectSlotSetChecked(QtObjectPtr ptr);
void QAction_DisconnectSlotSetChecked(QtObjectPtr ptr);
void QAction_SetChecked_Bool(QtObjectPtr ptr, int bo);
void QAction_ConnectSlotSetDisabled(QtObjectPtr ptr);
void QAction_DisconnectSlotSetDisabled(QtObjectPtr ptr);
void QAction_SetDisabled_Bool(QtObjectPtr ptr, int bo);
void QAction_ConnectSlotSetEnabled(QtObjectPtr ptr);
void QAction_DisconnectSlotSetEnabled(QtObjectPtr ptr);
void QAction_SetEnabled_Bool(QtObjectPtr ptr, int bo);
void QAction_ConnectSlotSetVisible(QtObjectPtr ptr);
void QAction_DisconnectSlotSetVisible(QtObjectPtr ptr);
void QAction_SetVisible_Bool(QtObjectPtr ptr, int bo);
void QAction_ConnectSlotToggle(QtObjectPtr ptr);
void QAction_DisconnectSlotToggle(QtObjectPtr ptr);
void QAction_Toggle(QtObjectPtr ptr);
void QAction_ConnectSlotTrigger(QtObjectPtr ptr);
void QAction_DisconnectSlotTrigger(QtObjectPtr ptr);
void QAction_Trigger(QtObjectPtr ptr);
//Signals
void QAction_ConnectSignalChanged(QtObjectPtr ptr);
void QAction_DisconnectSignalChanged(QtObjectPtr ptr);
void QAction_ConnectSignalHovered(QtObjectPtr ptr);
void QAction_DisconnectSignalHovered(QtObjectPtr ptr);
void QAction_ConnectSignalToggled(QtObjectPtr ptr);
void QAction_DisconnectSignalToggled(QtObjectPtr ptr);
void QAction_ConnectSignalTriggered(QtObjectPtr ptr);
void QAction_DisconnectSignalTriggered(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
