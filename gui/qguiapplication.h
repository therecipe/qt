#ifdef __cplusplus
extern "C" {
#endif

char* QGuiApplication_QGuiApplication_ApplicationDisplayName();
int QGuiApplication_QGuiApplication_ApplicationState();
int QGuiApplication_IsSavingSession(void* ptr);
int QGuiApplication_IsSessionRestored(void* ptr);
int QGuiApplication_QGuiApplication_LayoutDirection();
void* QGuiApplication_QGuiApplication_OverrideCursor();
char* QGuiApplication_QGuiApplication_PlatformName();
int QGuiApplication_QGuiApplication_QueryKeyboardModifiers();
int QGuiApplication_QGuiApplication_QuitOnLastWindowClosed();
void QGuiApplication_QGuiApplication_RestoreOverrideCursor();
char* QGuiApplication_SessionId(void* ptr);
char* QGuiApplication_SessionKey(void* ptr);
void QGuiApplication_QGuiApplication_SetApplicationDisplayName(char* name);
void QGuiApplication_QGuiApplication_SetLayoutDirection(int direction);
void QGuiApplication_QGuiApplication_SetOverrideCursor(void* cursor);
void QGuiApplication_QGuiApplication_SetQuitOnLastWindowClosed(int quit);
void QGuiApplication_QGuiApplication_SetWindowIcon(void* icon);
void* QGuiApplication_NewQGuiApplication(int argc, char* argv);
void QGuiApplication_ConnectApplicationStateChanged(void* ptr);
void QGuiApplication_DisconnectApplicationStateChanged(void* ptr);
void QGuiApplication_QGuiApplication_ChangeOverrideCursor(void* cursor);
void* QGuiApplication_QGuiApplication_Clipboard();
int QGuiApplication_QGuiApplication_DesktopSettingsAware();
double QGuiApplication_DevicePixelRatio(void* ptr);
int QGuiApplication_QGuiApplication_Exec();
void* QGuiApplication_QGuiApplication_FocusObject();
void QGuiApplication_ConnectFocusObjectChanged(void* ptr);
void QGuiApplication_DisconnectFocusObjectChanged(void* ptr);
void* QGuiApplication_QGuiApplication_FocusWindow();
void QGuiApplication_ConnectFocusWindowChanged(void* ptr);
void QGuiApplication_DisconnectFocusWindowChanged(void* ptr);
void QGuiApplication_ConnectFontDatabaseChanged(void* ptr);
void QGuiApplication_DisconnectFontDatabaseChanged(void* ptr);
void* QGuiApplication_QGuiApplication_InputMethod();
int QGuiApplication_QGuiApplication_IsLeftToRight();
int QGuiApplication_QGuiApplication_IsRightToLeft();
int QGuiApplication_QGuiApplication_KeyboardModifiers();
void QGuiApplication_ConnectLastWindowClosed(void* ptr);
void QGuiApplication_DisconnectLastWindowClosed(void* ptr);
void QGuiApplication_ConnectLayoutDirectionChanged(void* ptr);
void QGuiApplication_DisconnectLayoutDirectionChanged(void* ptr);
void* QGuiApplication_QGuiApplication_ModalWindow();
int QGuiApplication_QGuiApplication_MouseButtons();
int QGuiApplication_Notify(void* ptr, void* object, void* event);
void* QGuiApplication_QGuiApplication_PrimaryScreen();
void QGuiApplication_ConnectScreenAdded(void* ptr);
void QGuiApplication_DisconnectScreenAdded(void* ptr);
void QGuiApplication_ConnectScreenRemoved(void* ptr);
void QGuiApplication_DisconnectScreenRemoved(void* ptr);
void QGuiApplication_QGuiApplication_SetDesktopSettingsAware(int on);
void QGuiApplication_QGuiApplication_SetFont(void* font);
void QGuiApplication_QGuiApplication_SetPalette(void* pal);
void* QGuiApplication_QGuiApplication_StyleHints();
void QGuiApplication_QGuiApplication_Sync();
void* QGuiApplication_QGuiApplication_TopLevelAt(void* pos);
void QGuiApplication_DestroyQGuiApplication(void* ptr);

#ifdef __cplusplus
}
#endif