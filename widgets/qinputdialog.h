#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QInputDialog_CancelButtonText(QtObjectPtr ptr);
char* QInputDialog_ComboBoxItems(QtObjectPtr ptr);
int QInputDialog_DoubleDecimals(QtObjectPtr ptr);
int QInputDialog_InputMode(QtObjectPtr ptr);
int QInputDialog_IntMaximum(QtObjectPtr ptr);
int QInputDialog_IntMinimum(QtObjectPtr ptr);
int QInputDialog_IntStep(QtObjectPtr ptr);
int QInputDialog_IntValue(QtObjectPtr ptr);
int QInputDialog_IsComboBoxEditable(QtObjectPtr ptr);
char* QInputDialog_LabelText(QtObjectPtr ptr);
char* QInputDialog_OkButtonText(QtObjectPtr ptr);
int QInputDialog_Options(QtObjectPtr ptr);
void QInputDialog_SetCancelButtonText(QtObjectPtr ptr, char* text);
void QInputDialog_SetComboBoxEditable(QtObjectPtr ptr, int editable);
void QInputDialog_SetComboBoxItems(QtObjectPtr ptr, char* items);
void QInputDialog_SetDoubleDecimals(QtObjectPtr ptr, int decimals);
void QInputDialog_SetInputMode(QtObjectPtr ptr, int mode);
void QInputDialog_SetIntMaximum(QtObjectPtr ptr, int max);
void QInputDialog_SetIntMinimum(QtObjectPtr ptr, int min);
void QInputDialog_SetIntStep(QtObjectPtr ptr, int step);
void QInputDialog_SetIntValue(QtObjectPtr ptr, int value);
void QInputDialog_SetLabelText(QtObjectPtr ptr, char* text);
void QInputDialog_SetOkButtonText(QtObjectPtr ptr, char* text);
void QInputDialog_SetOptions(QtObjectPtr ptr, int options);
void QInputDialog_SetTextEchoMode(QtObjectPtr ptr, int mode);
void QInputDialog_SetTextValue(QtObjectPtr ptr, char* text);
int QInputDialog_TextEchoMode(QtObjectPtr ptr);
char* QInputDialog_TextValue(QtObjectPtr ptr);
QtObjectPtr QInputDialog_NewQInputDialog(QtObjectPtr parent, int flags);
void QInputDialog_Done(QtObjectPtr ptr, int result);
int QInputDialog_QInputDialog_GetInt(QtObjectPtr parent, char* title, char* label, int value, int min, int max, int step, int ok, int flags);
char* QInputDialog_QInputDialog_GetItem(QtObjectPtr parent, char* title, char* label, char* items, int current, int editable, int ok, int flags, int inputMethodHints);
char* QInputDialog_QInputDialog_GetMultiLineText(QtObjectPtr parent, char* title, char* label, char* text, int ok, int flags, int inputMethodHints);
char* QInputDialog_QInputDialog_GetText(QtObjectPtr parent, char* title, char* label, int mode, char* text, int ok, int flags, int inputMethodHints);
void QInputDialog_ConnectIntValueChanged(QtObjectPtr ptr);
void QInputDialog_DisconnectIntValueChanged(QtObjectPtr ptr);
void QInputDialog_ConnectIntValueSelected(QtObjectPtr ptr);
void QInputDialog_DisconnectIntValueSelected(QtObjectPtr ptr);
void QInputDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member);
void QInputDialog_SetIntRange(QtObjectPtr ptr, int min, int max);
void QInputDialog_SetOption(QtObjectPtr ptr, int option, int on);
void QInputDialog_SetVisible(QtObjectPtr ptr, int visible);
int QInputDialog_TestOption(QtObjectPtr ptr, int option);
void QInputDialog_ConnectTextValueChanged(QtObjectPtr ptr);
void QInputDialog_DisconnectTextValueChanged(QtObjectPtr ptr);
void QInputDialog_ConnectTextValueSelected(QtObjectPtr ptr);
void QInputDialog_DisconnectTextValueSelected(QtObjectPtr ptr);
void QInputDialog_DestroyQInputDialog(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif