#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMenuBar_IsDefaultUp(QtObjectPtr ptr);
int QMenuBar_IsNativeMenuBar(QtObjectPtr ptr);
void QMenuBar_SetCornerWidget(QtObjectPtr ptr, QtObjectPtr widget, int corner);
void QMenuBar_SetDefaultUp(QtObjectPtr ptr, int v);
void QMenuBar_SetNativeMenuBar(QtObjectPtr ptr, int nativeMenuBar);
QtObjectPtr QMenuBar_NewQMenuBar(QtObjectPtr parent);
QtObjectPtr QMenuBar_ActionAt(QtObjectPtr ptr, QtObjectPtr pt);
QtObjectPtr QMenuBar_ActiveAction(QtObjectPtr ptr);
QtObjectPtr QMenuBar_AddAction(QtObjectPtr ptr, char* text);
QtObjectPtr QMenuBar_AddAction2(QtObjectPtr ptr, char* text, QtObjectPtr receiver, char* member);
QtObjectPtr QMenuBar_AddMenu(QtObjectPtr ptr, QtObjectPtr menu);
QtObjectPtr QMenuBar_AddMenu3(QtObjectPtr ptr, QtObjectPtr icon, char* title);
QtObjectPtr QMenuBar_AddMenu2(QtObjectPtr ptr, char* title);
QtObjectPtr QMenuBar_AddSeparator(QtObjectPtr ptr);
void QMenuBar_Clear(QtObjectPtr ptr);
QtObjectPtr QMenuBar_CornerWidget(QtObjectPtr ptr, int corner);
int QMenuBar_HeightForWidth(QtObjectPtr ptr, int v);
void QMenuBar_ConnectHovered(QtObjectPtr ptr);
void QMenuBar_DisconnectHovered(QtObjectPtr ptr);
QtObjectPtr QMenuBar_InsertMenu(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr menu);
QtObjectPtr QMenuBar_InsertSeparator(QtObjectPtr ptr, QtObjectPtr before);
void QMenuBar_SetActiveAction(QtObjectPtr ptr, QtObjectPtr act);
void QMenuBar_SetVisible(QtObjectPtr ptr, int visible);
void QMenuBar_ConnectTriggered(QtObjectPtr ptr);
void QMenuBar_DisconnectTriggered(QtObjectPtr ptr);
void QMenuBar_DestroyQMenuBar(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif