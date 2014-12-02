#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
void QAbstractButton_Destroy(QtObjectPtr ptr);
int QAbstractButton_AutoExclusive(QtObjectPtr ptr);
int QAbstractButton_AutoRepeat(QtObjectPtr ptr);
int QAbstractButton_AutoRepeatDelay(QtObjectPtr ptr);
int QAbstractButton_AutoRepeatInterval(QtObjectPtr ptr);
int QAbstractButton_IsCheckable(QtObjectPtr ptr);
int QAbstractButton_IsChecked(QtObjectPtr ptr);
int QAbstractButton_IsDown(QtObjectPtr ptr);
void QAbstractButton_SetAutoExclusive_Bool(QtObjectPtr ptr, int autoExclusive);
void QAbstractButton_SetAutoRepeat_Bool(QtObjectPtr ptr, int autoRepeat);
void QAbstractButton_SetAutoRepeatDelay_Int(QtObjectPtr ptr, int autoRepeatDelay);
void QAbstractButton_SetAutoRepeatInterval_Int(QtObjectPtr ptr, int autoRepeatInterval);
void QAbstractButton_SetCheckable_Bool(QtObjectPtr ptr, int checkable);
void QAbstractButton_SetDown_Bool(QtObjectPtr ptr, int down);
void QAbstractButton_SetText_String(QtObjectPtr ptr, char* text);
char* QAbstractButton_Text(QtObjectPtr ptr);
//Public Slots
void QAbstractButton_ConnectSlotAnimateClick(QtObjectPtr ptr);
void QAbstractButton_DisconnectSlotAnimateClick(QtObjectPtr ptr);
void QAbstractButton_AnimateClick_Int(QtObjectPtr ptr, int msec);
void QAbstractButton_ConnectSlotClick(QtObjectPtr ptr);
void QAbstractButton_DisconnectSlotClick(QtObjectPtr ptr);
void QAbstractButton_Click(QtObjectPtr ptr);
void QAbstractButton_ConnectSlotSetChecked(QtObjectPtr ptr);
void QAbstractButton_DisconnectSlotSetChecked(QtObjectPtr ptr);
void QAbstractButton_SetChecked_Bool(QtObjectPtr ptr, int checked);
void QAbstractButton_ConnectSlotToggle(QtObjectPtr ptr);
void QAbstractButton_DisconnectSlotToggle(QtObjectPtr ptr);
void QAbstractButton_Toggle(QtObjectPtr ptr);
//Signals
void QAbstractButton_ConnectSignalClicked(QtObjectPtr ptr);
void QAbstractButton_DisconnectSignalClicked(QtObjectPtr ptr);
void QAbstractButton_ConnectSignalPressed(QtObjectPtr ptr);
void QAbstractButton_DisconnectSignalPressed(QtObjectPtr ptr);
void QAbstractButton_ConnectSignalReleased(QtObjectPtr ptr);
void QAbstractButton_DisconnectSignalReleased(QtObjectPtr ptr);
void QAbstractButton_ConnectSignalToggled(QtObjectPtr ptr);
void QAbstractButton_DisconnectSignalToggled(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
