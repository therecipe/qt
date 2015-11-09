#ifdef __cplusplus
extern "C" {
#endif

char* QInputDialog_CancelButtonText(void* ptr);
char* QInputDialog_ComboBoxItems(void* ptr);
int QInputDialog_DoubleDecimals(void* ptr);
int QInputDialog_InputMode(void* ptr);
int QInputDialog_IntMaximum(void* ptr);
int QInputDialog_IntMinimum(void* ptr);
int QInputDialog_IntStep(void* ptr);
int QInputDialog_IntValue(void* ptr);
int QInputDialog_IsComboBoxEditable(void* ptr);
char* QInputDialog_LabelText(void* ptr);
char* QInputDialog_OkButtonText(void* ptr);
int QInputDialog_Options(void* ptr);
void QInputDialog_SetCancelButtonText(void* ptr, char* text);
void QInputDialog_SetComboBoxEditable(void* ptr, int editable);
void QInputDialog_SetComboBoxItems(void* ptr, char* items);
void QInputDialog_SetDoubleDecimals(void* ptr, int decimals);
void QInputDialog_SetInputMode(void* ptr, int mode);
void QInputDialog_SetIntMaximum(void* ptr, int max);
void QInputDialog_SetIntMinimum(void* ptr, int min);
void QInputDialog_SetIntStep(void* ptr, int step);
void QInputDialog_SetIntValue(void* ptr, int value);
void QInputDialog_SetLabelText(void* ptr, char* text);
void QInputDialog_SetOkButtonText(void* ptr, char* text);
void QInputDialog_SetOptions(void* ptr, int options);
void QInputDialog_SetTextEchoMode(void* ptr, int mode);
void QInputDialog_SetTextValue(void* ptr, char* text);
int QInputDialog_TextEchoMode(void* ptr);
char* QInputDialog_TextValue(void* ptr);
void* QInputDialog_NewQInputDialog(void* parent, int flags);
void QInputDialog_Done(void* ptr, int result);
int QInputDialog_QInputDialog_GetInt(void* parent, char* title, char* label, int value, int min, int max, int step, int ok, int flags);
char* QInputDialog_QInputDialog_GetItem(void* parent, char* title, char* label, char* items, int current, int editable, int ok, int flags, int inputMethodHints);
char* QInputDialog_QInputDialog_GetMultiLineText(void* parent, char* title, char* label, char* text, int ok, int flags, int inputMethodHints);
char* QInputDialog_QInputDialog_GetText(void* parent, char* title, char* label, int mode, char* text, int ok, int flags, int inputMethodHints);
void QInputDialog_ConnectIntValueChanged(void* ptr);
void QInputDialog_DisconnectIntValueChanged(void* ptr);
void QInputDialog_ConnectIntValueSelected(void* ptr);
void QInputDialog_DisconnectIntValueSelected(void* ptr);
void QInputDialog_Open(void* ptr, void* receiver, char* member);
void QInputDialog_SetIntRange(void* ptr, int min, int max);
void QInputDialog_SetOption(void* ptr, int option, int on);
void QInputDialog_SetVisible(void* ptr, int visible);
int QInputDialog_TestOption(void* ptr, int option);
void QInputDialog_ConnectTextValueChanged(void* ptr);
void QInputDialog_DisconnectTextValueChanged(void* ptr);
void QInputDialog_ConnectTextValueSelected(void* ptr);
void QInputDialog_DisconnectTextValueSelected(void* ptr);
void QInputDialog_DestroyQInputDialog(void* ptr);

#ifdef __cplusplus
}
#endif