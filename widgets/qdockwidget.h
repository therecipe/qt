#ifdef __cplusplus
extern "C" {
#endif

int QDockWidget_AllowedAreas(void* ptr);
int QDockWidget_Features(void* ptr);
void QDockWidget_SetAllowedAreas(void* ptr, int areas);
void QDockWidget_SetFeatures(void* ptr, int features);
void QDockWidget_SetFloating(void* ptr, int floating);
void* QDockWidget_NewQDockWidget2(void* parent, int flags);
void* QDockWidget_NewQDockWidget(char* title, void* parent, int flags);
void QDockWidget_ConnectAllowedAreasChanged(void* ptr);
void QDockWidget_DisconnectAllowedAreasChanged(void* ptr);
void QDockWidget_ConnectDockLocationChanged(void* ptr);
void QDockWidget_DisconnectDockLocationChanged(void* ptr);
void QDockWidget_ConnectFeaturesChanged(void* ptr);
void QDockWidget_DisconnectFeaturesChanged(void* ptr);
int QDockWidget_IsAreaAllowed(void* ptr, int area);
int QDockWidget_IsFloating(void* ptr);
void QDockWidget_SetTitleBarWidget(void* ptr, void* widget);
void QDockWidget_SetWidget(void* ptr, void* widget);
void* QDockWidget_TitleBarWidget(void* ptr);
void* QDockWidget_ToggleViewAction(void* ptr);
void QDockWidget_ConnectTopLevelChanged(void* ptr);
void QDockWidget_DisconnectTopLevelChanged(void* ptr);
void QDockWidget_ConnectVisibilityChanged(void* ptr);
void QDockWidget_DisconnectVisibilityChanged(void* ptr);
void* QDockWidget_Widget(void* ptr);
void QDockWidget_DestroyQDockWidget(void* ptr);

#ifdef __cplusplus
}
#endif