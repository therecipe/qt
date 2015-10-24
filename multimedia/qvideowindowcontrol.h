#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QVideoWindowControl_AspectRatioMode(QtObjectPtr ptr);
int QVideoWindowControl_Brightness(QtObjectPtr ptr);
void QVideoWindowControl_ConnectBrightnessChanged(QtObjectPtr ptr);
void QVideoWindowControl_DisconnectBrightnessChanged(QtObjectPtr ptr);
int QVideoWindowControl_Contrast(QtObjectPtr ptr);
void QVideoWindowControl_ConnectContrastChanged(QtObjectPtr ptr);
void QVideoWindowControl_DisconnectContrastChanged(QtObjectPtr ptr);
void QVideoWindowControl_ConnectFullScreenChanged(QtObjectPtr ptr);
void QVideoWindowControl_DisconnectFullScreenChanged(QtObjectPtr ptr);
int QVideoWindowControl_Hue(QtObjectPtr ptr);
void QVideoWindowControl_ConnectHueChanged(QtObjectPtr ptr);
void QVideoWindowControl_DisconnectHueChanged(QtObjectPtr ptr);
int QVideoWindowControl_IsFullScreen(QtObjectPtr ptr);
void QVideoWindowControl_ConnectNativeSizeChanged(QtObjectPtr ptr);
void QVideoWindowControl_DisconnectNativeSizeChanged(QtObjectPtr ptr);
void QVideoWindowControl_Repaint(QtObjectPtr ptr);
int QVideoWindowControl_Saturation(QtObjectPtr ptr);
void QVideoWindowControl_ConnectSaturationChanged(QtObjectPtr ptr);
void QVideoWindowControl_DisconnectSaturationChanged(QtObjectPtr ptr);
void QVideoWindowControl_SetAspectRatioMode(QtObjectPtr ptr, int mode);
void QVideoWindowControl_SetBrightness(QtObjectPtr ptr, int brightness);
void QVideoWindowControl_SetContrast(QtObjectPtr ptr, int contrast);
void QVideoWindowControl_SetDisplayRect(QtObjectPtr ptr, QtObjectPtr rect);
void QVideoWindowControl_SetFullScreen(QtObjectPtr ptr, int fullScreen);
void QVideoWindowControl_SetHue(QtObjectPtr ptr, int hue);
void QVideoWindowControl_SetSaturation(QtObjectPtr ptr, int saturation);
void QVideoWindowControl_DestroyQVideoWindowControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif