#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QApplication_QApplication_Alert(QtObjectPtr widget, int msec);
int QApplication_AutoMaximizeThreshold(QtObjectPtr ptr);
int QApplication_AutoSipEnabled(QtObjectPtr ptr);
void QApplication_QApplication_Beep();
int QApplication_QApplication_CursorFlashTime();
int QApplication_QApplication_DoubleClickInterval();
int QApplication_QApplication_IsEffectEnabled(int effect);
int QApplication_QApplication_KeyboardInputInterval();
void QApplication_QApplication_SetActiveWindow(QtObjectPtr active);
void QApplication_SetAutoMaximizeThreshold(QtObjectPtr ptr, int threshold);
void QApplication_SetAutoSipEnabled(QtObjectPtr ptr, int enabled);
void QApplication_QApplication_SetCursorFlashTime(int v);
void QApplication_QApplication_SetDoubleClickInterval(int v);
void QApplication_QApplication_SetEffectEnabled(int effect, int enable);
void QApplication_QApplication_SetGlobalStrut(QtObjectPtr v);
void QApplication_QApplication_SetKeyboardInputInterval(int v);
void QApplication_QApplication_SetStartDragDistance(int l);
void QApplication_QApplication_SetStartDragTime(int ms);
void QApplication_SetStyleSheet(QtObjectPtr ptr, char* sheet);
void QApplication_QApplication_SetWheelScrollLines(int v);
void QApplication_QApplication_SetWindowIcon(QtObjectPtr icon);
int QApplication_QApplication_StartDragDistance();
int QApplication_QApplication_StartDragTime();
char* QApplication_StyleSheet(QtObjectPtr ptr);
QtObjectPtr QApplication_QApplication_TopLevelAt(QtObjectPtr point);
int QApplication_QApplication_WheelScrollLines();
QtObjectPtr QApplication_QApplication_WidgetAt(QtObjectPtr point);
QtObjectPtr QApplication_NewQApplication(int argc, char* argv);
void QApplication_QApplication_AboutQt();
QtObjectPtr QApplication_QApplication_ActiveModalWidget();
QtObjectPtr QApplication_QApplication_ActivePopupWidget();
QtObjectPtr QApplication_QApplication_ActiveWindow();
void QApplication_QApplication_CloseAllWindows();
int QApplication_QApplication_ColorSpec();
QtObjectPtr QApplication_QApplication_Desktop();
int QApplication_QApplication_Exec();
void QApplication_ConnectFocusChanged(QtObjectPtr ptr);
void QApplication_DisconnectFocusChanged(QtObjectPtr ptr);
QtObjectPtr QApplication_QApplication_FocusWidget();
int QApplication_Notify(QtObjectPtr ptr, QtObjectPtr receiver, QtObjectPtr e);
void QApplication_QApplication_SetColorSpec(int spec);
void QApplication_QApplication_SetFont(QtObjectPtr font, char* className);
void QApplication_QApplication_SetPalette(QtObjectPtr palette, char* className);
QtObjectPtr QApplication_QApplication_SetStyle2(char* style);
void QApplication_QApplication_SetStyle(QtObjectPtr style);
QtObjectPtr QApplication_QApplication_Style();
QtObjectPtr QApplication_QApplication_TopLevelAt2(int x, int y);
QtObjectPtr QApplication_QApplication_WidgetAt2(int x, int y);
void QApplication_DestroyQApplication(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif