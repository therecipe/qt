#ifdef __cplusplus
extern "C" {
#endif

int QTabBar_AutoHide(void* ptr);
int QTabBar_ChangeCurrentOnDrag(void* ptr);
int QTabBar_Count(void* ptr);
int QTabBar_CurrentIndex(void* ptr);
int QTabBar_DocumentMode(void* ptr);
int QTabBar_DrawBase(void* ptr);
int QTabBar_ElideMode(void* ptr);
int QTabBar_Expanding(void* ptr);
int QTabBar_IsMovable(void* ptr);
int QTabBar_SelectionBehaviorOnRemove(void* ptr);
void QTabBar_SetAutoHide(void* ptr, int hide);
void QTabBar_SetChangeCurrentOnDrag(void* ptr, int change);
void QTabBar_SetCurrentIndex(void* ptr, int index);
void QTabBar_SetDocumentMode(void* ptr, int set);
void QTabBar_SetDrawBase(void* ptr, int drawTheBase);
void QTabBar_SetElideMode(void* ptr, int v);
void QTabBar_SetExpanding(void* ptr, int enabled);
void QTabBar_SetIconSize(void* ptr, void* size);
void QTabBar_SetMovable(void* ptr, int movable);
void QTabBar_SetSelectionBehaviorOnRemove(void* ptr, int behavior);
void QTabBar_SetShape(void* ptr, int shape);
void QTabBar_SetTabsClosable(void* ptr, int closable);
void QTabBar_SetUsesScrollButtons(void* ptr, int useButtons);
int QTabBar_Shape(void* ptr);
int QTabBar_TabsClosable(void* ptr);
int QTabBar_UsesScrollButtons(void* ptr);
void* QTabBar_NewQTabBar(void* parent);
int QTabBar_AddTab2(void* ptr, void* icon, char* text);
int QTabBar_AddTab(void* ptr, char* text);
void QTabBar_ConnectCurrentChanged(void* ptr);
void QTabBar_DisconnectCurrentChanged(void* ptr);
int QTabBar_InsertTab2(void* ptr, int index, void* icon, char* text);
int QTabBar_InsertTab(void* ptr, int index, char* text);
int QTabBar_IsTabEnabled(void* ptr, int index);
void QTabBar_MoveTab(void* ptr, int from, int to);
void QTabBar_RemoveTab(void* ptr, int index);
void QTabBar_SetTabButton(void* ptr, int index, int position, void* widget);
void QTabBar_SetTabData(void* ptr, int index, void* data);
void QTabBar_SetTabEnabled(void* ptr, int index, int enabled);
void QTabBar_SetTabIcon(void* ptr, int index, void* icon);
void QTabBar_SetTabText(void* ptr, int index, char* text);
void QTabBar_SetTabTextColor(void* ptr, int index, void* color);
void QTabBar_SetTabToolTip(void* ptr, int index, char* tip);
void QTabBar_SetTabWhatsThis(void* ptr, int index, char* text);
int QTabBar_TabAt(void* ptr, void* position);
void QTabBar_ConnectTabBarClicked(void* ptr);
void QTabBar_DisconnectTabBarClicked(void* ptr);
void QTabBar_ConnectTabBarDoubleClicked(void* ptr);
void QTabBar_DisconnectTabBarDoubleClicked(void* ptr);
void* QTabBar_TabButton(void* ptr, int index, int position);
void QTabBar_ConnectTabCloseRequested(void* ptr);
void QTabBar_DisconnectTabCloseRequested(void* ptr);
void* QTabBar_TabData(void* ptr, int index);
void QTabBar_ConnectTabMoved(void* ptr);
void QTabBar_DisconnectTabMoved(void* ptr);
char* QTabBar_TabText(void* ptr, int index);
void* QTabBar_TabTextColor(void* ptr, int index);
char* QTabBar_TabToolTip(void* ptr, int index);
char* QTabBar_TabWhatsThis(void* ptr, int index);
void QTabBar_DestroyQTabBar(void* ptr);

#ifdef __cplusplus
}
#endif