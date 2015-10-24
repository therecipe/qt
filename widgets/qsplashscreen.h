#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSplashScreen_NewQSplashScreen2(QtObjectPtr parent, QtObjectPtr pixmap, int f);
QtObjectPtr QSplashScreen_NewQSplashScreen(QtObjectPtr pixmap, int f);
void QSplashScreen_ClearMessage(QtObjectPtr ptr);
void QSplashScreen_Finish(QtObjectPtr ptr, QtObjectPtr mainWin);
char* QSplashScreen_Message(QtObjectPtr ptr);
void QSplashScreen_ConnectMessageChanged(QtObjectPtr ptr);
void QSplashScreen_DisconnectMessageChanged(QtObjectPtr ptr);
void QSplashScreen_Repaint(QtObjectPtr ptr);
void QSplashScreen_SetPixmap(QtObjectPtr ptr, QtObjectPtr pixmap);
void QSplashScreen_ShowMessage(QtObjectPtr ptr, char* message, int alignment, QtObjectPtr color);
void QSplashScreen_DestroyQSplashScreen(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif