#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDesktopWidget_IsVirtualDesktop(QtObjectPtr ptr);
int QDesktopWidget_PrimaryScreen(QtObjectPtr ptr);
QtObjectPtr QDesktopWidget_Screen(QtObjectPtr ptr, int screen);
int QDesktopWidget_ScreenNumber2(QtObjectPtr ptr, QtObjectPtr point);
int QDesktopWidget_ScreenNumber(QtObjectPtr ptr, QtObjectPtr widget);
void QDesktopWidget_ConnectResized(QtObjectPtr ptr);
void QDesktopWidget_DisconnectResized(QtObjectPtr ptr);
int QDesktopWidget_ScreenCount(QtObjectPtr ptr);
void QDesktopWidget_ConnectScreenCountChanged(QtObjectPtr ptr);
void QDesktopWidget_DisconnectScreenCountChanged(QtObjectPtr ptr);
void QDesktopWidget_ConnectWorkAreaResized(QtObjectPtr ptr);
void QDesktopWidget_DisconnectWorkAreaResized(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif