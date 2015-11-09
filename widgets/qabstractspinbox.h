#ifdef __cplusplus
extern "C" {
#endif

int QAbstractSpinBox_Alignment(void* ptr);
int QAbstractSpinBox_ButtonSymbols(void* ptr);
int QAbstractSpinBox_CorrectionMode(void* ptr);
int QAbstractSpinBox_HasAcceptableInput(void* ptr);
int QAbstractSpinBox_HasFrame(void* ptr);
int QAbstractSpinBox_IsAccelerated(void* ptr);
int QAbstractSpinBox_IsGroupSeparatorShown(void* ptr);
int QAbstractSpinBox_IsReadOnly(void* ptr);
int QAbstractSpinBox_KeyboardTracking(void* ptr);
void QAbstractSpinBox_SetAccelerated(void* ptr, int on);
void QAbstractSpinBox_SetAlignment(void* ptr, int flag);
void QAbstractSpinBox_SetButtonSymbols(void* ptr, int bs);
void QAbstractSpinBox_SetCorrectionMode(void* ptr, int cm);
void QAbstractSpinBox_SetFrame(void* ptr, int v);
void QAbstractSpinBox_SetGroupSeparatorShown(void* ptr, int shown);
void QAbstractSpinBox_SetKeyboardTracking(void* ptr, int kt);
void QAbstractSpinBox_SetReadOnly(void* ptr, int r);
void QAbstractSpinBox_SetSpecialValueText(void* ptr, char* txt);
void QAbstractSpinBox_SetWrapping(void* ptr, int w);
char* QAbstractSpinBox_SpecialValueText(void* ptr);
char* QAbstractSpinBox_Text(void* ptr);
int QAbstractSpinBox_Wrapping(void* ptr);
void* QAbstractSpinBox_NewQAbstractSpinBox(void* parent);
void QAbstractSpinBox_Clear(void* ptr);
void QAbstractSpinBox_ConnectEditingFinished(void* ptr);
void QAbstractSpinBox_DisconnectEditingFinished(void* ptr);
int QAbstractSpinBox_Event(void* ptr, void* event);
void* QAbstractSpinBox_InputMethodQuery(void* ptr, int query);
void QAbstractSpinBox_InterpretText(void* ptr);
void QAbstractSpinBox_SelectAll(void* ptr);
void QAbstractSpinBox_StepBy(void* ptr, int steps);
void QAbstractSpinBox_StepDown(void* ptr);
void QAbstractSpinBox_StepUp(void* ptr);
void QAbstractSpinBox_DestroyQAbstractSpinBox(void* ptr);

#ifdef __cplusplus
}
#endif