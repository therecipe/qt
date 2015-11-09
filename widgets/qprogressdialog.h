#ifdef __cplusplus
extern "C" {
#endif

int QProgressDialog_AutoClose(void* ptr);
int QProgressDialog_AutoReset(void* ptr);
char* QProgressDialog_LabelText(void* ptr);
int QProgressDialog_Maximum(void* ptr);
int QProgressDialog_Minimum(void* ptr);
int QProgressDialog_MinimumDuration(void* ptr);
void QProgressDialog_SetAutoClose(void* ptr, int close);
void QProgressDialog_SetAutoReset(void* ptr, int reset);
void QProgressDialog_SetLabelText(void* ptr, char* text);
void QProgressDialog_SetMaximum(void* ptr, int maximum);
void QProgressDialog_SetMinimum(void* ptr, int minimum);
void QProgressDialog_SetMinimumDuration(void* ptr, int ms);
void QProgressDialog_SetValue(void* ptr, int progress);
int QProgressDialog_Value(void* ptr);
int QProgressDialog_WasCanceled(void* ptr);
void* QProgressDialog_NewQProgressDialog(void* parent, int f);
void* QProgressDialog_NewQProgressDialog2(char* labelText, char* cancelButtonText, int minimum, int maximum, void* parent, int f);
void QProgressDialog_Cancel(void* ptr);
void QProgressDialog_ConnectCanceled(void* ptr);
void QProgressDialog_DisconnectCanceled(void* ptr);
void QProgressDialog_Open(void* ptr, void* receiver, char* member);
void QProgressDialog_Reset(void* ptr);
void QProgressDialog_SetBar(void* ptr, void* bar);
void QProgressDialog_SetCancelButton(void* ptr, void* cancelButton);
void QProgressDialog_SetCancelButtonText(void* ptr, char* cancelButtonText);
void QProgressDialog_SetLabel(void* ptr, void* label);
void QProgressDialog_SetRange(void* ptr, int minimum, int maximum);
void QProgressDialog_DestroyQProgressDialog(void* ptr);

#ifdef __cplusplus
}
#endif