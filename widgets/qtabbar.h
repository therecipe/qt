#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTabBar_AutoHide(QtObjectPtr ptr);
int QTabBar_ChangeCurrentOnDrag(QtObjectPtr ptr);
int QTabBar_Count(QtObjectPtr ptr);
int QTabBar_CurrentIndex(QtObjectPtr ptr);
int QTabBar_DocumentMode(QtObjectPtr ptr);
int QTabBar_DrawBase(QtObjectPtr ptr);
int QTabBar_ElideMode(QtObjectPtr ptr);
int QTabBar_Expanding(QtObjectPtr ptr);
int QTabBar_IsMovable(QtObjectPtr ptr);
int QTabBar_SelectionBehaviorOnRemove(QtObjectPtr ptr);
void QTabBar_SetAutoHide(QtObjectPtr ptr, int hide);
void QTabBar_SetChangeCurrentOnDrag(QtObjectPtr ptr, int change);
void QTabBar_SetCurrentIndex(QtObjectPtr ptr, int index);
void QTabBar_SetDocumentMode(QtObjectPtr ptr, int set);
void QTabBar_SetDrawBase(QtObjectPtr ptr, int drawTheBase);
void QTabBar_SetElideMode(QtObjectPtr ptr, int v);
void QTabBar_SetExpanding(QtObjectPtr ptr, int enabled);
void QTabBar_SetIconSize(QtObjectPtr ptr, QtObjectPtr size);
void QTabBar_SetMovable(QtObjectPtr ptr, int movable);
void QTabBar_SetSelectionBehaviorOnRemove(QtObjectPtr ptr, int behavior);
void QTabBar_SetShape(QtObjectPtr ptr, int shape);
void QTabBar_SetTabsClosable(QtObjectPtr ptr, int closable);
void QTabBar_SetUsesScrollButtons(QtObjectPtr ptr, int useButtons);
int QTabBar_Shape(QtObjectPtr ptr);
int QTabBar_TabsClosable(QtObjectPtr ptr);
int QTabBar_UsesScrollButtons(QtObjectPtr ptr);
QtObjectPtr QTabBar_NewQTabBar(QtObjectPtr parent);
int QTabBar_AddTab2(QtObjectPtr ptr, QtObjectPtr icon, char* text);
int QTabBar_AddTab(QtObjectPtr ptr, char* text);
void QTabBar_ConnectCurrentChanged(QtObjectPtr ptr);
void QTabBar_DisconnectCurrentChanged(QtObjectPtr ptr);
int QTabBar_InsertTab2(QtObjectPtr ptr, int index, QtObjectPtr icon, char* text);
int QTabBar_InsertTab(QtObjectPtr ptr, int index, char* text);
int QTabBar_IsTabEnabled(QtObjectPtr ptr, int index);
void QTabBar_MoveTab(QtObjectPtr ptr, int from, int to);
void QTabBar_RemoveTab(QtObjectPtr ptr, int index);
void QTabBar_SetTabButton(QtObjectPtr ptr, int index, int position, QtObjectPtr widget);
void QTabBar_SetTabData(QtObjectPtr ptr, int index, char* data);
void QTabBar_SetTabEnabled(QtObjectPtr ptr, int index, int enabled);
void QTabBar_SetTabIcon(QtObjectPtr ptr, int index, QtObjectPtr icon);
void QTabBar_SetTabText(QtObjectPtr ptr, int index, char* text);
void QTabBar_SetTabTextColor(QtObjectPtr ptr, int index, QtObjectPtr color);
void QTabBar_SetTabToolTip(QtObjectPtr ptr, int index, char* tip);
void QTabBar_SetTabWhatsThis(QtObjectPtr ptr, int index, char* text);
int QTabBar_TabAt(QtObjectPtr ptr, QtObjectPtr position);
void QTabBar_ConnectTabBarClicked(QtObjectPtr ptr);
void QTabBar_DisconnectTabBarClicked(QtObjectPtr ptr);
void QTabBar_ConnectTabBarDoubleClicked(QtObjectPtr ptr);
void QTabBar_DisconnectTabBarDoubleClicked(QtObjectPtr ptr);
QtObjectPtr QTabBar_TabButton(QtObjectPtr ptr, int index, int position);
void QTabBar_ConnectTabCloseRequested(QtObjectPtr ptr);
void QTabBar_DisconnectTabCloseRequested(QtObjectPtr ptr);
char* QTabBar_TabData(QtObjectPtr ptr, int index);
void QTabBar_ConnectTabMoved(QtObjectPtr ptr);
void QTabBar_DisconnectTabMoved(QtObjectPtr ptr);
char* QTabBar_TabText(QtObjectPtr ptr, int index);
char* QTabBar_TabToolTip(QtObjectPtr ptr, int index);
char* QTabBar_TabWhatsThis(QtObjectPtr ptr, int index);
void QTabBar_DestroyQTabBar(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif