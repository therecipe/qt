#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QProgressDialog_AutoClose(QtObjectPtr ptr);
int QProgressDialog_AutoReset(QtObjectPtr ptr);
char* QProgressDialog_LabelText(QtObjectPtr ptr);
int QProgressDialog_Maximum(QtObjectPtr ptr);
int QProgressDialog_Minimum(QtObjectPtr ptr);
int QProgressDialog_MinimumDuration(QtObjectPtr ptr);
void QProgressDialog_SetAutoClose(QtObjectPtr ptr, int close);
void QProgressDialog_SetAutoReset(QtObjectPtr ptr, int reset);
void QProgressDialog_SetLabelText(QtObjectPtr ptr, char* text);
void QProgressDialog_SetMaximum(QtObjectPtr ptr, int maximum);
void QProgressDialog_SetMinimum(QtObjectPtr ptr, int minimum);
void QProgressDialog_SetMinimumDuration(QtObjectPtr ptr, int ms);
void QProgressDialog_SetValue(QtObjectPtr ptr, int progress);
int QProgressDialog_Value(QtObjectPtr ptr);
int QProgressDialog_WasCanceled(QtObjectPtr ptr);
QtObjectPtr QProgressDialog_NewQProgressDialog(QtObjectPtr parent, int f);
QtObjectPtr QProgressDialog_NewQProgressDialog2(char* labelText, char* cancelButtonText, int minimum, int maximum, QtObjectPtr parent, int f);
void QProgressDialog_Cancel(QtObjectPtr ptr);
void QProgressDialog_ConnectCanceled(QtObjectPtr ptr);
void QProgressDialog_DisconnectCanceled(QtObjectPtr ptr);
void QProgressDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member);
void QProgressDialog_Reset(QtObjectPtr ptr);
void QProgressDialog_SetBar(QtObjectPtr ptr, QtObjectPtr bar);
void QProgressDialog_SetCancelButton(QtObjectPtr ptr, QtObjectPtr cancelButton);
void QProgressDialog_SetCancelButtonText(QtObjectPtr ptr, char* cancelButtonText);
void QProgressDialog_SetLabel(QtObjectPtr ptr, QtObjectPtr label);
void QProgressDialog_SetRange(QtObjectPtr ptr, int minimum, int maximum);
void QProgressDialog_DestroyQProgressDialog(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif