#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QApplication_New_Int_String(int argc, char* argv);
void QApplication_Destroy(QtObjectPtr ptr);
char* QApplication_StyleSheet(QtObjectPtr ptr);
//Public Slots
void QApplication_ConnectSlotAutoSipEnabled(QtObjectPtr ptr);
void QApplication_DisconnectSlotAutoSipEnabled(QtObjectPtr ptr);
void QApplication_AutoSipEnabled(QtObjectPtr ptr);
//Static Public Members
QtObjectPtr QApplication_ActiveModalWidget();
QtObjectPtr QApplication_ActivePopupWidget();
QtObjectPtr QApplication_ActiveWindow();
void QApplication_Alert_QWidget_Int(QtObjectPtr widget, int msec);
int QApplication_ColorSpec();
int QApplication_CursorFlashTime();
int QApplication_DoubleClickInterval();
int QApplication_Exec();
QtObjectPtr QApplication_FocusWidget();
int QApplication_IsEffectEnabled_UIEffect(int effect);
int QApplication_KeyboardInputInterval();
void QApplication_SetActiveWindow_QWidget(QtObjectPtr active);
void QApplication_SetColorSpec_Int(int spec);
void QApplication_SetCursorFlashTime_Int(int flashTime);
void QApplication_SetDoubleClickInterval_Int(int doubleClickInterval);
void QApplication_SetEffectEnabled_UIEffect_Bool(int effect, int enable);
void QApplication_SetKeyboardInputInterval_Int(int keyboardInputInterval);
void QApplication_SetStartDragDistance_Int(int l);
void QApplication_SetStartDragTime_Int(int ms);
void QApplication_SetWheelScrollLines_Int(int wheelScrollLines);
int QApplication_StartDragDistance();
int QApplication_StartDragTime();
QtObjectPtr QApplication_TopLevelAt_Int_Int(int x, int y);
int QApplication_WheelScrollLines();
QtObjectPtr QApplication_WidgetAt_Int_Int(int x, int y);

#ifdef __cplusplus
}
#endif
