#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QStatusBar_New_QWidget(QtObjectPtr parent);
void QStatusBar_Destroy(QtObjectPtr ptr);
void QStatusBar_AddPermanentWidget_QWidget_Int(QtObjectPtr ptr, QtObjectPtr widget, int stretch);
void QStatusBar_AddWidget_QWidget_Int(QtObjectPtr ptr, QtObjectPtr widget, int stretch);
char* QStatusBar_CurrentMessage(QtObjectPtr ptr);
int QStatusBar_InsertPermanentWidget_Int_QWidget_Int(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch);
int QStatusBar_InsertWidget_Int_QWidget_Int(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch);
int QStatusBar_IsSizeGripEnabled(QtObjectPtr ptr);
void QStatusBar_RemoveWidget_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QStatusBar_SetSizeGripEnabled_Bool(QtObjectPtr ptr, int enabled);
//Public Slots
void QStatusBar_ConnectSlotClearMessage(QtObjectPtr ptr);
void QStatusBar_DisconnectSlotClearMessage(QtObjectPtr ptr);
void QStatusBar_ClearMessage(QtObjectPtr ptr);
//Signals
void QStatusBar_ConnectSignalMessageChanged(QtObjectPtr ptr);
void QStatusBar_DisconnectSignalMessageChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
