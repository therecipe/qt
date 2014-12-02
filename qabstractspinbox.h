#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QAbstractSpinBox_UpDownArrows();
int QAbstractSpinBox_PlusMinus();
int QAbstractSpinBox_NoButtons();
int QAbstractSpinBox_CorrectToPreviousValue();
int QAbstractSpinBox_CorrectToNearestValue();
int QAbstractSpinBox_StepNone();
int QAbstractSpinBox_StepUpEnabled();
int QAbstractSpinBox_StepDownEnabled();
//Public Functions
QtObjectPtr QAbstractSpinBox_New_QWidget(QtObjectPtr parent);
void QAbstractSpinBox_Destroy(QtObjectPtr ptr);
int QAbstractSpinBox_Alignment(QtObjectPtr ptr);
int QAbstractSpinBox_ButtonSymbols(QtObjectPtr ptr);
int QAbstractSpinBox_HasAcceptableInput(QtObjectPtr ptr);
int QAbstractSpinBox_HasFrame(QtObjectPtr ptr);
void QAbstractSpinBox_InterpretText(QtObjectPtr ptr);
int QAbstractSpinBox_IsAccelerated(QtObjectPtr ptr);
int QAbstractSpinBox_IsGroupSeparatorShown(QtObjectPtr ptr);
int QAbstractSpinBox_IsReadOnly(QtObjectPtr ptr);
int QAbstractSpinBox_KeyboardTracking(QtObjectPtr ptr);
void QAbstractSpinBox_SetAccelerated_Bool(QtObjectPtr ptr, int on);
void QAbstractSpinBox_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int flag);
void QAbstractSpinBox_SetButtonSymbols_ButtonSymbols(QtObjectPtr ptr, int bs);
void QAbstractSpinBox_SetFrame_Bool(QtObjectPtr ptr, int frame);
void QAbstractSpinBox_SetGroupSeparatorShown_Bool(QtObjectPtr ptr, int shown);
void QAbstractSpinBox_SetKeyboardTracking_Bool(QtObjectPtr ptr, int kt);
void QAbstractSpinBox_SetReadOnly_Bool(QtObjectPtr ptr, int r);
void QAbstractSpinBox_SetSpecialValueText_String(QtObjectPtr ptr, char* txt);
void QAbstractSpinBox_SetWrapping_Bool(QtObjectPtr ptr, int w);
char* QAbstractSpinBox_SpecialValueText(QtObjectPtr ptr);
void QAbstractSpinBox_StepBy_Int(QtObjectPtr ptr, int steps);
char* QAbstractSpinBox_Text(QtObjectPtr ptr);
int QAbstractSpinBox_Wrapping(QtObjectPtr ptr);
//Public Slots
void QAbstractSpinBox_ConnectSlotClear(QtObjectPtr ptr);
void QAbstractSpinBox_DisconnectSlotClear(QtObjectPtr ptr);
void QAbstractSpinBox_Clear(QtObjectPtr ptr);
void QAbstractSpinBox_ConnectSlotSelectAll(QtObjectPtr ptr);
void QAbstractSpinBox_DisconnectSlotSelectAll(QtObjectPtr ptr);
void QAbstractSpinBox_SelectAll(QtObjectPtr ptr);
void QAbstractSpinBox_ConnectSlotStepDown(QtObjectPtr ptr);
void QAbstractSpinBox_DisconnectSlotStepDown(QtObjectPtr ptr);
void QAbstractSpinBox_StepDown(QtObjectPtr ptr);
void QAbstractSpinBox_ConnectSlotStepUp(QtObjectPtr ptr);
void QAbstractSpinBox_DisconnectSlotStepUp(QtObjectPtr ptr);
void QAbstractSpinBox_StepUp(QtObjectPtr ptr);
//Signals
void QAbstractSpinBox_ConnectSignalEditingFinished(QtObjectPtr ptr);
void QAbstractSpinBox_DisconnectSignalEditingFinished(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
