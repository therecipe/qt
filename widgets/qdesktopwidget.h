#ifdef __cplusplus
extern "C" {
#endif

int QDesktopWidget_IsVirtualDesktop(void* ptr);
int QDesktopWidget_PrimaryScreen(void* ptr);
void* QDesktopWidget_Screen(void* ptr, int screen);
int QDesktopWidget_ScreenNumber2(void* ptr, void* point);
int QDesktopWidget_ScreenNumber(void* ptr, void* widget);
void QDesktopWidget_ConnectResized(void* ptr);
void QDesktopWidget_DisconnectResized(void* ptr);
int QDesktopWidget_ScreenCount(void* ptr);
void QDesktopWidget_ConnectScreenCountChanged(void* ptr);
void QDesktopWidget_DisconnectScreenCountChanged(void* ptr);
void QDesktopWidget_ConnectWorkAreaResized(void* ptr);
void QDesktopWidget_DisconnectWorkAreaResized(void* ptr);

#ifdef __cplusplus
}
#endif