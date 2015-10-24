#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAbstractSpinBox_Alignment(QtObjectPtr ptr);
int QAbstractSpinBox_ButtonSymbols(QtObjectPtr ptr);
int QAbstractSpinBox_CorrectionMode(QtObjectPtr ptr);
int QAbstractSpinBox_HasAcceptableInput(QtObjectPtr ptr);
int QAbstractSpinBox_HasFrame(QtObjectPtr ptr);
int QAbstractSpinBox_IsAccelerated(QtObjectPtr ptr);
int QAbstractSpinBox_IsGroupSeparatorShown(QtObjectPtr ptr);
int QAbstractSpinBox_IsReadOnly(QtObjectPtr ptr);
int QAbstractSpinBox_KeyboardTracking(QtObjectPtr ptr);
void QAbstractSpinBox_SetAccelerated(QtObjectPtr ptr, int on);
void QAbstractSpinBox_SetAlignment(QtObjectPtr ptr, int flag);
void QAbstractSpinBox_SetButtonSymbols(QtObjectPtr ptr, int bs);
void QAbstractSpinBox_SetCorrectionMode(QtObjectPtr ptr, int cm);
void QAbstractSpinBox_SetFrame(QtObjectPtr ptr, int v);
void QAbstractSpinBox_SetGroupSeparatorShown(QtObjectPtr ptr, int shown);
void QAbstractSpinBox_SetKeyboardTracking(QtObjectPtr ptr, int kt);
void QAbstractSpinBox_SetReadOnly(QtObjectPtr ptr, int r);
void QAbstractSpinBox_SetSpecialValueText(QtObjectPtr ptr, char* txt);
void QAbstractSpinBox_SetWrapping(QtObjectPtr ptr, int w);
char* QAbstractSpinBox_SpecialValueText(QtObjectPtr ptr);
char* QAbstractSpinBox_Text(QtObjectPtr ptr);
int QAbstractSpinBox_Wrapping(QtObjectPtr ptr);
QtObjectPtr QAbstractSpinBox_NewQAbstractSpinBox(QtObjectPtr parent);
void QAbstractSpinBox_Clear(QtObjectPtr ptr);
void QAbstractSpinBox_ConnectEditingFinished(QtObjectPtr ptr);
void QAbstractSpinBox_DisconnectEditingFinished(QtObjectPtr ptr);
int QAbstractSpinBox_Event(QtObjectPtr ptr, QtObjectPtr event);
char* QAbstractSpinBox_InputMethodQuery(QtObjectPtr ptr, int query);
void QAbstractSpinBox_InterpretText(QtObjectPtr ptr);
void QAbstractSpinBox_SelectAll(QtObjectPtr ptr);
void QAbstractSpinBox_StepBy(QtObjectPtr ptr, int steps);
void QAbstractSpinBox_StepDown(QtObjectPtr ptr);
void QAbstractSpinBox_StepUp(QtObjectPtr ptr);
void QAbstractSpinBox_DestroyQAbstractSpinBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif