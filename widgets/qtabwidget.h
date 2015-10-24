#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTabWidget_AddTab2(QtObjectPtr ptr, QtObjectPtr page, QtObjectPtr icon, char* label);
int QTabWidget_AddTab(QtObjectPtr ptr, QtObjectPtr page, char* label);
int QTabWidget_Count(QtObjectPtr ptr);
int QTabWidget_CurrentIndex(QtObjectPtr ptr);
int QTabWidget_DocumentMode(QtObjectPtr ptr);
int QTabWidget_ElideMode(QtObjectPtr ptr);
int QTabWidget_InsertTab2(QtObjectPtr ptr, int index, QtObjectPtr page, QtObjectPtr icon, char* label);
int QTabWidget_InsertTab(QtObjectPtr ptr, int index, QtObjectPtr page, char* label);
int QTabWidget_IsMovable(QtObjectPtr ptr);
void QTabWidget_SetCornerWidget(QtObjectPtr ptr, QtObjectPtr widget, int corner);
void QTabWidget_SetCurrentIndex(QtObjectPtr ptr, int index);
void QTabWidget_SetDocumentMode(QtObjectPtr ptr, int set);
void QTabWidget_SetElideMode(QtObjectPtr ptr, int v);
void QTabWidget_SetIconSize(QtObjectPtr ptr, QtObjectPtr size);
void QTabWidget_SetMovable(QtObjectPtr ptr, int movable);
void QTabWidget_SetTabBarAutoHide(QtObjectPtr ptr, int enabled);
void QTabWidget_SetTabPosition(QtObjectPtr ptr, int v);
void QTabWidget_SetTabShape(QtObjectPtr ptr, int s);
void QTabWidget_SetTabsClosable(QtObjectPtr ptr, int closeable);
void QTabWidget_SetUsesScrollButtons(QtObjectPtr ptr, int useButtons);
int QTabWidget_TabBarAutoHide(QtObjectPtr ptr);
int QTabWidget_TabPosition(QtObjectPtr ptr);
int QTabWidget_TabShape(QtObjectPtr ptr);
int QTabWidget_TabsClosable(QtObjectPtr ptr);
int QTabWidget_UsesScrollButtons(QtObjectPtr ptr);
QtObjectPtr QTabWidget_NewQTabWidget(QtObjectPtr parent);
void QTabWidget_Clear(QtObjectPtr ptr);
QtObjectPtr QTabWidget_CornerWidget(QtObjectPtr ptr, int corner);
void QTabWidget_ConnectCurrentChanged(QtObjectPtr ptr);
void QTabWidget_DisconnectCurrentChanged(QtObjectPtr ptr);
QtObjectPtr QTabWidget_CurrentWidget(QtObjectPtr ptr);
int QTabWidget_HasHeightForWidth(QtObjectPtr ptr);
int QTabWidget_HeightForWidth(QtObjectPtr ptr, int width);
int QTabWidget_IndexOf(QtObjectPtr ptr, QtObjectPtr w);
int QTabWidget_IsTabEnabled(QtObjectPtr ptr, int index);
void QTabWidget_RemoveTab(QtObjectPtr ptr, int index);
void QTabWidget_SetCurrentWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QTabWidget_SetTabEnabled(QtObjectPtr ptr, int index, int enable);
void QTabWidget_SetTabIcon(QtObjectPtr ptr, int index, QtObjectPtr icon);
void QTabWidget_SetTabText(QtObjectPtr ptr, int index, char* label);
void QTabWidget_SetTabToolTip(QtObjectPtr ptr, int index, char* tip);
void QTabWidget_SetTabWhatsThis(QtObjectPtr ptr, int index, char* text);
QtObjectPtr QTabWidget_TabBar(QtObjectPtr ptr);
void QTabWidget_ConnectTabBarClicked(QtObjectPtr ptr);
void QTabWidget_DisconnectTabBarClicked(QtObjectPtr ptr);
void QTabWidget_ConnectTabBarDoubleClicked(QtObjectPtr ptr);
void QTabWidget_DisconnectTabBarDoubleClicked(QtObjectPtr ptr);
void QTabWidget_ConnectTabCloseRequested(QtObjectPtr ptr);
void QTabWidget_DisconnectTabCloseRequested(QtObjectPtr ptr);
char* QTabWidget_TabText(QtObjectPtr ptr, int index);
char* QTabWidget_TabToolTip(QtObjectPtr ptr, int index);
char* QTabWidget_TabWhatsThis(QtObjectPtr ptr, int index);
QtObjectPtr QTabWidget_Widget(QtObjectPtr ptr, int index);
void QTabWidget_DestroyQTabWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif