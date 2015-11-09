#ifdef __cplusplus
extern "C" {
#endif

int QTabWidget_AddTab2(void* ptr, void* page, void* icon, char* label);
int QTabWidget_AddTab(void* ptr, void* page, char* label);
int QTabWidget_Count(void* ptr);
int QTabWidget_CurrentIndex(void* ptr);
int QTabWidget_DocumentMode(void* ptr);
int QTabWidget_ElideMode(void* ptr);
int QTabWidget_InsertTab2(void* ptr, int index, void* page, void* icon, char* label);
int QTabWidget_InsertTab(void* ptr, int index, void* page, char* label);
int QTabWidget_IsMovable(void* ptr);
void QTabWidget_SetCornerWidget(void* ptr, void* widget, int corner);
void QTabWidget_SetCurrentIndex(void* ptr, int index);
void QTabWidget_SetDocumentMode(void* ptr, int set);
void QTabWidget_SetElideMode(void* ptr, int v);
void QTabWidget_SetIconSize(void* ptr, void* size);
void QTabWidget_SetMovable(void* ptr, int movable);
void QTabWidget_SetTabBarAutoHide(void* ptr, int enabled);
void QTabWidget_SetTabPosition(void* ptr, int v);
void QTabWidget_SetTabShape(void* ptr, int s);
void QTabWidget_SetTabsClosable(void* ptr, int closeable);
void QTabWidget_SetUsesScrollButtons(void* ptr, int useButtons);
int QTabWidget_TabBarAutoHide(void* ptr);
int QTabWidget_TabPosition(void* ptr);
int QTabWidget_TabShape(void* ptr);
int QTabWidget_TabsClosable(void* ptr);
int QTabWidget_UsesScrollButtons(void* ptr);
void* QTabWidget_NewQTabWidget(void* parent);
void QTabWidget_Clear(void* ptr);
void* QTabWidget_CornerWidget(void* ptr, int corner);
void QTabWidget_ConnectCurrentChanged(void* ptr);
void QTabWidget_DisconnectCurrentChanged(void* ptr);
void* QTabWidget_CurrentWidget(void* ptr);
int QTabWidget_HasHeightForWidth(void* ptr);
int QTabWidget_HeightForWidth(void* ptr, int width);
int QTabWidget_IndexOf(void* ptr, void* w);
int QTabWidget_IsTabEnabled(void* ptr, int index);
void QTabWidget_RemoveTab(void* ptr, int index);
void QTabWidget_SetCurrentWidget(void* ptr, void* widget);
void QTabWidget_SetTabEnabled(void* ptr, int index, int enable);
void QTabWidget_SetTabIcon(void* ptr, int index, void* icon);
void QTabWidget_SetTabText(void* ptr, int index, char* label);
void QTabWidget_SetTabToolTip(void* ptr, int index, char* tip);
void QTabWidget_SetTabWhatsThis(void* ptr, int index, char* text);
void* QTabWidget_TabBar(void* ptr);
void QTabWidget_ConnectTabBarClicked(void* ptr);
void QTabWidget_DisconnectTabBarClicked(void* ptr);
void QTabWidget_ConnectTabBarDoubleClicked(void* ptr);
void QTabWidget_DisconnectTabBarDoubleClicked(void* ptr);
void QTabWidget_ConnectTabCloseRequested(void* ptr);
void QTabWidget_DisconnectTabCloseRequested(void* ptr);
char* QTabWidget_TabText(void* ptr, int index);
char* QTabWidget_TabToolTip(void* ptr, int index);
char* QTabWidget_TabWhatsThis(void* ptr, int index);
void* QTabWidget_Widget(void* ptr, int index);
void QTabWidget_DestroyQTabWidget(void* ptr);

#ifdef __cplusplus
}
#endif