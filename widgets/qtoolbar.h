#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QToolBar_AllowedAreas(QtObjectPtr ptr);
int QToolBar_IsFloatable(QtObjectPtr ptr);
int QToolBar_IsFloating(QtObjectPtr ptr);
int QToolBar_IsMovable(QtObjectPtr ptr);
int QToolBar_Orientation(QtObjectPtr ptr);
void QToolBar_SetAllowedAreas(QtObjectPtr ptr, int areas);
void QToolBar_SetFloatable(QtObjectPtr ptr, int floatable);
void QToolBar_SetIconSize(QtObjectPtr ptr, QtObjectPtr iconSize);
void QToolBar_SetMovable(QtObjectPtr ptr, int movable);
void QToolBar_SetOrientation(QtObjectPtr ptr, int orientation);
void QToolBar_SetToolButtonStyle(QtObjectPtr ptr, int toolButtonStyle);
int QToolBar_ToolButtonStyle(QtObjectPtr ptr);
QtObjectPtr QToolBar_NewQToolBar2(QtObjectPtr parent);
QtObjectPtr QToolBar_NewQToolBar(char* title, QtObjectPtr parent);
QtObjectPtr QToolBar_ActionAt(QtObjectPtr ptr, QtObjectPtr p);
QtObjectPtr QToolBar_ActionAt2(QtObjectPtr ptr, int x, int y);
void QToolBar_ConnectActionTriggered(QtObjectPtr ptr);
void QToolBar_DisconnectActionTriggered(QtObjectPtr ptr);
QtObjectPtr QToolBar_AddAction2(QtObjectPtr ptr, QtObjectPtr icon, char* text);
QtObjectPtr QToolBar_AddAction4(QtObjectPtr ptr, QtObjectPtr icon, char* text, QtObjectPtr receiver, char* member);
QtObjectPtr QToolBar_AddAction(QtObjectPtr ptr, char* text);
QtObjectPtr QToolBar_AddAction3(QtObjectPtr ptr, char* text, QtObjectPtr receiver, char* member);
QtObjectPtr QToolBar_AddSeparator(QtObjectPtr ptr);
QtObjectPtr QToolBar_AddWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QToolBar_ConnectAllowedAreasChanged(QtObjectPtr ptr);
void QToolBar_DisconnectAllowedAreasChanged(QtObjectPtr ptr);
void QToolBar_Clear(QtObjectPtr ptr);
QtObjectPtr QToolBar_InsertSeparator(QtObjectPtr ptr, QtObjectPtr before);
QtObjectPtr QToolBar_InsertWidget(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr widget);
int QToolBar_IsAreaAllowed(QtObjectPtr ptr, int area);
void QToolBar_ConnectMovableChanged(QtObjectPtr ptr);
void QToolBar_DisconnectMovableChanged(QtObjectPtr ptr);
void QToolBar_ConnectOrientationChanged(QtObjectPtr ptr);
void QToolBar_DisconnectOrientationChanged(QtObjectPtr ptr);
QtObjectPtr QToolBar_ToggleViewAction(QtObjectPtr ptr);
void QToolBar_ConnectToolButtonStyleChanged(QtObjectPtr ptr);
void QToolBar_DisconnectToolButtonStyleChanged(QtObjectPtr ptr);
void QToolBar_ConnectTopLevelChanged(QtObjectPtr ptr);
void QToolBar_DisconnectTopLevelChanged(QtObjectPtr ptr);
void QToolBar_ConnectVisibilityChanged(QtObjectPtr ptr);
void QToolBar_DisconnectVisibilityChanged(QtObjectPtr ptr);
QtObjectPtr QToolBar_WidgetForAction(QtObjectPtr ptr, QtObjectPtr action);
void QToolBar_DestroyQToolBar(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif