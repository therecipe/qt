#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QGuiApplication_New_Int_String(int argc, char* argv);
void QGuiApplication_Destroy(QtObjectPtr ptr);
int QGuiApplication_IsSavingSession(QtObjectPtr ptr);
int QGuiApplication_IsSessionRestored(QtObjectPtr ptr);
char* QGuiApplication_SessionId(QtObjectPtr ptr);
char* QGuiApplication_SessionKey(QtObjectPtr ptr);
//Signals
void QGuiApplication_ConnectSignalApplicationStateChanged(QtObjectPtr ptr);
void QGuiApplication_DisconnectSignalApplicationStateChanged(QtObjectPtr ptr);
void QGuiApplication_ConnectSignalFocusObjectChanged(QtObjectPtr ptr);
void QGuiApplication_DisconnectSignalFocusObjectChanged(QtObjectPtr ptr);
void QGuiApplication_ConnectSignalFontDatabaseChanged(QtObjectPtr ptr);
void QGuiApplication_DisconnectSignalFontDatabaseChanged(QtObjectPtr ptr);
void QGuiApplication_ConnectSignalLastWindowClosed(QtObjectPtr ptr);
void QGuiApplication_DisconnectSignalLastWindowClosed(QtObjectPtr ptr);
//Static Public Members
char* QGuiApplication_ApplicationDisplayName();
int QGuiApplication_ApplicationState();
int QGuiApplication_DesktopSettingsAware();
QtObjectPtr QGuiApplication_FocusObject();
int QGuiApplication_IsLeftToRight();
int QGuiApplication_IsRightToLeft();
int QGuiApplication_KeyboardModifiers();
int QGuiApplication_LayoutDirection();
int QGuiApplication_MouseButtons();
char* QGuiApplication_PlatformName();
int QGuiApplication_QueryKeyboardModifiers();
int QGuiApplication_QuitOnLastWindowClosed();
void QGuiApplication_SetApplicationDisplayName_String(char* name);
void QGuiApplication_SetDesktopSettingsAware_Bool(int on);
void QGuiApplication_SetLayoutDirection_LayoutDirection(int direction);
void QGuiApplication_SetQuitOnLastWindowClosed_Bool(int quit);

#ifdef __cplusplus
}
#endif
