#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QProgressDialog_New_QWidget_WindowType(QtObjectPtr parent, int f);
QtObjectPtr QProgressDialog_New_String_String_Int_Int_QWidget_WindowType(char* labelText, char* cancelButtonText, int minimum, int maximum, QtObjectPtr parent, int f);
void QProgressDialog_Destroy(QtObjectPtr ptr);
int QProgressDialog_AutoClose(QtObjectPtr ptr);
int QProgressDialog_AutoReset(QtObjectPtr ptr);
char* QProgressDialog_LabelText(QtObjectPtr ptr);
int QProgressDialog_Maximum(QtObjectPtr ptr);
int QProgressDialog_Minimum(QtObjectPtr ptr);
int QProgressDialog_MinimumDuration(QtObjectPtr ptr);
void QProgressDialog_SetAutoClose_Bool(QtObjectPtr ptr, int close);
void QProgressDialog_SetAutoReset_Bool(QtObjectPtr ptr, int reset);
void QProgressDialog_SetBar_QProgressBar(QtObjectPtr ptr, QtObjectPtr bar);
void QProgressDialog_SetCancelButton_QPushButton(QtObjectPtr ptr, QtObjectPtr cancelButton);
void QProgressDialog_SetLabel_QLabel(QtObjectPtr ptr, QtObjectPtr label);
int QProgressDialog_Value(QtObjectPtr ptr);
int QProgressDialog_WasCanceled(QtObjectPtr ptr);
//Public Slots
void QProgressDialog_ConnectSlotCancel(QtObjectPtr ptr);
void QProgressDialog_DisconnectSlotCancel(QtObjectPtr ptr);
void QProgressDialog_Cancel(QtObjectPtr ptr);
void QProgressDialog_ConnectSlotReset(QtObjectPtr ptr);
void QProgressDialog_DisconnectSlotReset(QtObjectPtr ptr);
void QProgressDialog_Reset(QtObjectPtr ptr);
void QProgressDialog_ConnectSlotSetMaximum(QtObjectPtr ptr);
void QProgressDialog_DisconnectSlotSetMaximum(QtObjectPtr ptr);
void QProgressDialog_SetMaximum_Int(QtObjectPtr ptr, int maximum);
void QProgressDialog_ConnectSlotSetMinimum(QtObjectPtr ptr);
void QProgressDialog_DisconnectSlotSetMinimum(QtObjectPtr ptr);
void QProgressDialog_SetMinimum_Int(QtObjectPtr ptr, int minimum);
void QProgressDialog_ConnectSlotSetMinimumDuration(QtObjectPtr ptr);
void QProgressDialog_DisconnectSlotSetMinimumDuration(QtObjectPtr ptr);
void QProgressDialog_SetMinimumDuration_Int(QtObjectPtr ptr, int ms);
void QProgressDialog_ConnectSlotSetRange(QtObjectPtr ptr);
void QProgressDialog_DisconnectSlotSetRange(QtObjectPtr ptr);
void QProgressDialog_SetRange_Int_Int(QtObjectPtr ptr, int minimum, int maximum);
void QProgressDialog_ConnectSlotSetValue(QtObjectPtr ptr);
void QProgressDialog_DisconnectSlotSetValue(QtObjectPtr ptr);
void QProgressDialog_SetValue_Int(QtObjectPtr ptr, int progress);
//Signals
void QProgressDialog_ConnectSignalCanceled(QtObjectPtr ptr);
void QProgressDialog_DisconnectSignalCanceled(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
