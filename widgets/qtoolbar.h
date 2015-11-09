#ifdef __cplusplus
extern "C" {
#endif

int QToolBar_AllowedAreas(void* ptr);
int QToolBar_IsFloatable(void* ptr);
int QToolBar_IsFloating(void* ptr);
int QToolBar_IsMovable(void* ptr);
int QToolBar_Orientation(void* ptr);
void QToolBar_SetAllowedAreas(void* ptr, int areas);
void QToolBar_SetFloatable(void* ptr, int floatable);
void QToolBar_SetIconSize(void* ptr, void* iconSize);
void QToolBar_SetMovable(void* ptr, int movable);
void QToolBar_SetOrientation(void* ptr, int orientation);
void QToolBar_SetToolButtonStyle(void* ptr, int toolButtonStyle);
int QToolBar_ToolButtonStyle(void* ptr);
void* QToolBar_NewQToolBar2(void* parent);
void* QToolBar_NewQToolBar(char* title, void* parent);
void* QToolBar_ActionAt(void* ptr, void* p);
void* QToolBar_ActionAt2(void* ptr, int x, int y);
void QToolBar_ConnectActionTriggered(void* ptr);
void QToolBar_DisconnectActionTriggered(void* ptr);
void* QToolBar_AddAction2(void* ptr, void* icon, char* text);
void* QToolBar_AddAction4(void* ptr, void* icon, char* text, void* receiver, char* member);
void* QToolBar_AddAction(void* ptr, char* text);
void* QToolBar_AddAction3(void* ptr, char* text, void* receiver, char* member);
void* QToolBar_AddSeparator(void* ptr);
void* QToolBar_AddWidget(void* ptr, void* widget);
void QToolBar_ConnectAllowedAreasChanged(void* ptr);
void QToolBar_DisconnectAllowedAreasChanged(void* ptr);
void QToolBar_Clear(void* ptr);
void* QToolBar_InsertSeparator(void* ptr, void* before);
void* QToolBar_InsertWidget(void* ptr, void* before, void* widget);
int QToolBar_IsAreaAllowed(void* ptr, int area);
void QToolBar_ConnectMovableChanged(void* ptr);
void QToolBar_DisconnectMovableChanged(void* ptr);
void QToolBar_ConnectOrientationChanged(void* ptr);
void QToolBar_DisconnectOrientationChanged(void* ptr);
void* QToolBar_ToggleViewAction(void* ptr);
void QToolBar_ConnectToolButtonStyleChanged(void* ptr);
void QToolBar_DisconnectToolButtonStyleChanged(void* ptr);
void QToolBar_ConnectTopLevelChanged(void* ptr);
void QToolBar_DisconnectTopLevelChanged(void* ptr);
void QToolBar_ConnectVisibilityChanged(void* ptr);
void QToolBar_DisconnectVisibilityChanged(void* ptr);
void* QToolBar_WidgetForAction(void* ptr, void* action);
void QToolBar_DestroyQToolBar(void* ptr);

#ifdef __cplusplus
}
#endif