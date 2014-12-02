#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QTabBar_New_QWidget(QtObjectPtr parent);
void QTabBar_Destroy(QtObjectPtr ptr);
int QTabBar_AddTab_String(QtObjectPtr ptr, char* text);
int QTabBar_Count(QtObjectPtr ptr);
int QTabBar_CurrentIndex(QtObjectPtr ptr);
int QTabBar_DocumentMode(QtObjectPtr ptr);
int QTabBar_DrawBase(QtObjectPtr ptr);
int QTabBar_ElideMode(QtObjectPtr ptr);
int QTabBar_Expanding(QtObjectPtr ptr);
int QTabBar_InsertTab_Int_String(QtObjectPtr ptr, int index, char* text);
int QTabBar_IsMovable(QtObjectPtr ptr);
int QTabBar_IsTabEnabled_Int(QtObjectPtr ptr, int index);
void QTabBar_MoveTab_Int_Int(QtObjectPtr ptr, int from, int to);
void QTabBar_RemoveTab_Int(QtObjectPtr ptr, int index);
void QTabBar_SetDocumentMode_Bool(QtObjectPtr ptr, int set);
void QTabBar_SetDrawBase_Bool(QtObjectPtr ptr, int drawTheBase);
void QTabBar_SetElideMode_TextElideMode(QtObjectPtr ptr, int TextElideMode);
void QTabBar_SetExpanding_Bool(QtObjectPtr ptr, int enabled);
void QTabBar_SetMovable_Bool(QtObjectPtr ptr, int movable);
void QTabBar_SetTabEnabled_Int_Bool(QtObjectPtr ptr, int index, int enabled);
void QTabBar_SetTabText_Int_String(QtObjectPtr ptr, int index, char* text);
void QTabBar_SetTabToolTip_Int_String(QtObjectPtr ptr, int index, char* tip);
void QTabBar_SetTabWhatsThis_Int_String(QtObjectPtr ptr, int index, char* text);
void QTabBar_SetTabsClosable_Bool(QtObjectPtr ptr, int closable);
void QTabBar_SetUsesScrollButtons_Bool(QtObjectPtr ptr, int useButtons);
char* QTabBar_TabText_Int(QtObjectPtr ptr, int index);
char* QTabBar_TabToolTip_Int(QtObjectPtr ptr, int index);
char* QTabBar_TabWhatsThis_Int(QtObjectPtr ptr, int index);
int QTabBar_TabsClosable(QtObjectPtr ptr);
int QTabBar_UsesScrollButtons(QtObjectPtr ptr);
//Public Slots
void QTabBar_ConnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QTabBar_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QTabBar_SetCurrentIndex_Int(QtObjectPtr ptr, int index);
//Signals
void QTabBar_ConnectSignalCurrentChanged(QtObjectPtr ptr);
void QTabBar_DisconnectSignalCurrentChanged(QtObjectPtr ptr);
void QTabBar_ConnectSignalTabBarClicked(QtObjectPtr ptr);
void QTabBar_DisconnectSignalTabBarClicked(QtObjectPtr ptr);
void QTabBar_ConnectSignalTabBarDoubleClicked(QtObjectPtr ptr);
void QTabBar_DisconnectSignalTabBarDoubleClicked(QtObjectPtr ptr);
void QTabBar_ConnectSignalTabCloseRequested(QtObjectPtr ptr);
void QTabBar_DisconnectSignalTabCloseRequested(QtObjectPtr ptr);
void QTabBar_ConnectSignalTabMoved(QtObjectPtr ptr);
void QTabBar_DisconnectSignalTabMoved(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
