#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QMenuBar_New_QWidget(QtObjectPtr parent);
void QMenuBar_Destroy(QtObjectPtr ptr);
QtObjectPtr QMenuBar_ActiveAction(QtObjectPtr ptr);
QtObjectPtr QMenuBar_AddAction_String(QtObjectPtr ptr, char* text);
QtObjectPtr QMenuBar_AddAction_String_QObject_String(QtObjectPtr ptr, char* text, QtObjectPtr receiver, char* member);
void QMenuBar_AddAction_QAction(QtObjectPtr ptr, QtObjectPtr action);
QtObjectPtr QMenuBar_AddMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu);
QtObjectPtr QMenuBar_AddMenu_String(QtObjectPtr ptr, char* title);
QtObjectPtr QMenuBar_AddSeparator(QtObjectPtr ptr);
void QMenuBar_Clear(QtObjectPtr ptr);
QtObjectPtr QMenuBar_CornerWidget_Corner(QtObjectPtr ptr, int corner);
QtObjectPtr QMenuBar_InsertMenu_QAction_QMenu(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr menu);
QtObjectPtr QMenuBar_InsertSeparator_QAction(QtObjectPtr ptr, QtObjectPtr before);
int QMenuBar_IsDefaultUp(QtObjectPtr ptr);
int QMenuBar_IsNativeMenuBar(QtObjectPtr ptr);
void QMenuBar_SetActiveAction_QAction(QtObjectPtr ptr, QtObjectPtr act);
void QMenuBar_SetCornerWidget_QWidget_Corner(QtObjectPtr ptr, QtObjectPtr widget, int corner);
void QMenuBar_SetDefaultUp_Bool(QtObjectPtr ptr, int defaultAction);
void QMenuBar_SetNativeMenuBar_Bool(QtObjectPtr ptr, int nativeMenuBar);
//Public Slots
void QMenuBar_ConnectSlotSetVisible(QtObjectPtr ptr);
void QMenuBar_DisconnectSlotSetVisible(QtObjectPtr ptr);
void QMenuBar_SetVisible_Bool(QtObjectPtr ptr, int visible);
//Signals
void QMenuBar_ConnectSignalHovered(QtObjectPtr ptr);
void QMenuBar_DisconnectSignalHovered(QtObjectPtr ptr);
void QMenuBar_ConnectSignalTriggered(QtObjectPtr ptr);
void QMenuBar_DisconnectSignalTriggered(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
