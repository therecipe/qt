#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QTabWidget_New_QWidget(QtObjectPtr parent);
void QTabWidget_Destroy(QtObjectPtr ptr);
int QTabWidget_AddTab_QWidget_String(QtObjectPtr ptr, QtObjectPtr page, char* label);
void QTabWidget_Clear(QtObjectPtr ptr);
QtObjectPtr QTabWidget_CornerWidget_Corner(QtObjectPtr ptr, int corner);
int QTabWidget_Count(QtObjectPtr ptr);
int QTabWidget_CurrentIndex(QtObjectPtr ptr);
QtObjectPtr QTabWidget_CurrentWidget(QtObjectPtr ptr);
int QTabWidget_DocumentMode(QtObjectPtr ptr);
int QTabWidget_ElideMode(QtObjectPtr ptr);
int QTabWidget_IndexOf_QWidget(QtObjectPtr ptr, QtObjectPtr w);
int QTabWidget_InsertTab_Int_QWidget_String(QtObjectPtr ptr, int index, QtObjectPtr page, char* label);
int QTabWidget_IsMovable(QtObjectPtr ptr);
int QTabWidget_IsTabEnabled_Int(QtObjectPtr ptr, int index);
void QTabWidget_RemoveTab_Int(QtObjectPtr ptr, int index);
void QTabWidget_SetCornerWidget_QWidget_Corner(QtObjectPtr ptr, QtObjectPtr widget, int corner);
void QTabWidget_SetDocumentMode_Bool(QtObjectPtr ptr, int set);
void QTabWidget_SetElideMode_TextElideMode(QtObjectPtr ptr, int TextElideMode);
void QTabWidget_SetMovable_Bool(QtObjectPtr ptr, int movable);
void QTabWidget_SetTabEnabled_Int_Bool(QtObjectPtr ptr, int index, int enable);
void QTabWidget_SetTabText_Int_String(QtObjectPtr ptr, int index, char* label);
void QTabWidget_SetTabToolTip_Int_String(QtObjectPtr ptr, int index, char* tip);
void QTabWidget_SetTabWhatsThis_Int_String(QtObjectPtr ptr, int index, char* text);
void QTabWidget_SetTabsClosable_Bool(QtObjectPtr ptr, int closeable);
void QTabWidget_SetUsesScrollButtons_Bool(QtObjectPtr ptr, int useButtons);
QtObjectPtr QTabWidget_TabBar(QtObjectPtr ptr);
char* QTabWidget_TabText_Int(QtObjectPtr ptr, int index);
char* QTabWidget_TabToolTip_Int(QtObjectPtr ptr, int index);
char* QTabWidget_TabWhatsThis_Int(QtObjectPtr ptr, int index);
int QTabWidget_TabsClosable(QtObjectPtr ptr);
int QTabWidget_UsesScrollButtons(QtObjectPtr ptr);
QtObjectPtr QTabWidget_Widget_Int(QtObjectPtr ptr, int index);
//Public Slots
void QTabWidget_ConnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QTabWidget_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QTabWidget_SetCurrentIndex_Int(QtObjectPtr ptr, int index);
//Signals
void QTabWidget_ConnectSignalCurrentChanged(QtObjectPtr ptr);
void QTabWidget_DisconnectSignalCurrentChanged(QtObjectPtr ptr);
void QTabWidget_ConnectSignalTabBarClicked(QtObjectPtr ptr);
void QTabWidget_DisconnectSignalTabBarClicked(QtObjectPtr ptr);
void QTabWidget_ConnectSignalTabBarDoubleClicked(QtObjectPtr ptr);
void QTabWidget_DisconnectSignalTabBarDoubleClicked(QtObjectPtr ptr);
void QTabWidget_ConnectSignalTabCloseRequested(QtObjectPtr ptr);
void QTabWidget_DisconnectSignalTabCloseRequested(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
