#ifdef __cplusplus
extern "C" {
#endif

void* QSplashScreen_NewQSplashScreen2(void* parent, void* pixmap, int f);
void* QSplashScreen_NewQSplashScreen(void* pixmap, int f);
void QSplashScreen_ClearMessage(void* ptr);
void QSplashScreen_Finish(void* ptr, void* mainWin);
char* QSplashScreen_Message(void* ptr);
void QSplashScreen_ConnectMessageChanged(void* ptr);
void QSplashScreen_DisconnectMessageChanged(void* ptr);
void QSplashScreen_Repaint(void* ptr);
void QSplashScreen_SetPixmap(void* ptr, void* pixmap);
void QSplashScreen_ShowMessage(void* ptr, char* message, int alignment, void* color);
void QSplashScreen_DestroyQSplashScreen(void* ptr);

#ifdef __cplusplus
}
#endif