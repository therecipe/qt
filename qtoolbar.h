#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QToolBar_New_String_QWidget(char* title, QtObjectPtr parent);
QtObjectPtr QToolBar_New_QWidget(QtObjectPtr parent);
void QToolBar_Destroy(QtObjectPtr ptr);
int QToolBar_AllowedAreas(QtObjectPtr ptr);
void QToolBar_Clear(QtObjectPtr ptr);
int QToolBar_IsAreaAllowed_ToolBarArea(QtObjectPtr ptr, int area);
int QToolBar_IsFloatable(QtObjectPtr ptr);
int QToolBar_IsFloating(QtObjectPtr ptr);
int QToolBar_IsMovable(QtObjectPtr ptr);
int QToolBar_Orientation(QtObjectPtr ptr);
void QToolBar_SetAllowedAreas_ToolBarArea(QtObjectPtr ptr, int areas);
void QToolBar_SetFloatable_Bool(QtObjectPtr ptr, int floatable);
void QToolBar_SetMovable_Bool(QtObjectPtr ptr, int movable);
void QToolBar_SetOrientation_Orientation(QtObjectPtr ptr, int orientation);
int QToolBar_ToolButtonStyle(QtObjectPtr ptr);
//Signals
void QToolBar_ConnectSignalActionTriggered(QtObjectPtr ptr);
void QToolBar_DisconnectSignalActionTriggered(QtObjectPtr ptr);
void QToolBar_ConnectSignalAllowedAreasChanged(QtObjectPtr ptr);
void QToolBar_DisconnectSignalAllowedAreasChanged(QtObjectPtr ptr);
void QToolBar_ConnectSignalMovableChanged(QtObjectPtr ptr);
void QToolBar_DisconnectSignalMovableChanged(QtObjectPtr ptr);
void QToolBar_ConnectSignalOrientationChanged(QtObjectPtr ptr);
void QToolBar_DisconnectSignalOrientationChanged(QtObjectPtr ptr);
void QToolBar_ConnectSignalToolButtonStyleChanged(QtObjectPtr ptr);
void QToolBar_DisconnectSignalToolButtonStyleChanged(QtObjectPtr ptr);
void QToolBar_ConnectSignalTopLevelChanged(QtObjectPtr ptr);
void QToolBar_DisconnectSignalTopLevelChanged(QtObjectPtr ptr);
void QToolBar_ConnectSignalVisibilityChanged(QtObjectPtr ptr);
void QToolBar_DisconnectSignalVisibilityChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
