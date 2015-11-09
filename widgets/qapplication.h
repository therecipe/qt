#ifdef __cplusplus
extern "C" {
#endif

void QApplication_QApplication_Alert(void* widget, int msec);
int QApplication_AutoMaximizeThreshold(void* ptr);
int QApplication_AutoSipEnabled(void* ptr);
void QApplication_QApplication_Beep();
int QApplication_QApplication_CursorFlashTime();
int QApplication_QApplication_DoubleClickInterval();
int QApplication_QApplication_IsEffectEnabled(int effect);
int QApplication_QApplication_KeyboardInputInterval();
void QApplication_QApplication_SetActiveWindow(void* active);
void QApplication_SetAutoMaximizeThreshold(void* ptr, int threshold);
void QApplication_SetAutoSipEnabled(void* ptr, int enabled);
void QApplication_QApplication_SetCursorFlashTime(int v);
void QApplication_QApplication_SetDoubleClickInterval(int v);
void QApplication_QApplication_SetEffectEnabled(int effect, int enable);
void QApplication_QApplication_SetGlobalStrut(void* v);
void QApplication_QApplication_SetKeyboardInputInterval(int v);
void QApplication_QApplication_SetStartDragDistance(int l);
void QApplication_QApplication_SetStartDragTime(int ms);
void QApplication_SetStyleSheet(void* ptr, char* sheet);
void QApplication_QApplication_SetWheelScrollLines(int v);
void QApplication_QApplication_SetWindowIcon(void* icon);
int QApplication_QApplication_StartDragDistance();
int QApplication_QApplication_StartDragTime();
char* QApplication_StyleSheet(void* ptr);
void* QApplication_QApplication_TopLevelAt(void* point);
int QApplication_QApplication_WheelScrollLines();
void* QApplication_QApplication_WidgetAt(void* point);
void* QApplication_NewQApplication(int argc, char* argv);
void QApplication_QApplication_AboutQt();
void* QApplication_QApplication_ActiveModalWidget();
void* QApplication_QApplication_ActivePopupWidget();
void* QApplication_QApplication_ActiveWindow();
void QApplication_QApplication_CloseAllWindows();
int QApplication_QApplication_ColorSpec();
void* QApplication_QApplication_Desktop();
int QApplication_QApplication_Exec();
void QApplication_ConnectFocusChanged(void* ptr);
void QApplication_DisconnectFocusChanged(void* ptr);
void* QApplication_QApplication_FocusWidget();
int QApplication_Notify(void* ptr, void* receiver, void* e);
void QApplication_QApplication_SetColorSpec(int spec);
void QApplication_QApplication_SetFont(void* font, char* className);
void QApplication_QApplication_SetPalette(void* palette, char* className);
void* QApplication_QApplication_SetStyle2(char* style);
void QApplication_QApplication_SetStyle(void* style);
void* QApplication_QApplication_Style();
void* QApplication_QApplication_TopLevelAt2(int x, int y);
void* QApplication_QApplication_WidgetAt2(int x, int y);
void QApplication_DestroyQApplication(void* ptr);

#ifdef __cplusplus
}
#endif