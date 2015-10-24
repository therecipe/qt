#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDockWidget_AllowedAreas(QtObjectPtr ptr);
int QDockWidget_Features(QtObjectPtr ptr);
void QDockWidget_SetAllowedAreas(QtObjectPtr ptr, int areas);
void QDockWidget_SetFeatures(QtObjectPtr ptr, int features);
void QDockWidget_SetFloating(QtObjectPtr ptr, int floating);
QtObjectPtr QDockWidget_NewQDockWidget2(QtObjectPtr parent, int flags);
QtObjectPtr QDockWidget_NewQDockWidget(char* title, QtObjectPtr parent, int flags);
void QDockWidget_ConnectAllowedAreasChanged(QtObjectPtr ptr);
void QDockWidget_DisconnectAllowedAreasChanged(QtObjectPtr ptr);
void QDockWidget_ConnectDockLocationChanged(QtObjectPtr ptr);
void QDockWidget_DisconnectDockLocationChanged(QtObjectPtr ptr);
void QDockWidget_ConnectFeaturesChanged(QtObjectPtr ptr);
void QDockWidget_DisconnectFeaturesChanged(QtObjectPtr ptr);
int QDockWidget_IsAreaAllowed(QtObjectPtr ptr, int area);
int QDockWidget_IsFloating(QtObjectPtr ptr);
void QDockWidget_SetTitleBarWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QDockWidget_SetWidget(QtObjectPtr ptr, QtObjectPtr widget);
QtObjectPtr QDockWidget_TitleBarWidget(QtObjectPtr ptr);
QtObjectPtr QDockWidget_ToggleViewAction(QtObjectPtr ptr);
void QDockWidget_ConnectTopLevelChanged(QtObjectPtr ptr);
void QDockWidget_DisconnectTopLevelChanged(QtObjectPtr ptr);
void QDockWidget_ConnectVisibilityChanged(QtObjectPtr ptr);
void QDockWidget_DisconnectVisibilityChanged(QtObjectPtr ptr);
QtObjectPtr QDockWidget_Widget(QtObjectPtr ptr);
void QDockWidget_DestroyQDockWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif