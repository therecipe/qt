#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QVideoWidgetControl_AspectRatioMode(QtObjectPtr ptr);
int QVideoWidgetControl_Brightness(QtObjectPtr ptr);
void QVideoWidgetControl_ConnectBrightnessChanged(QtObjectPtr ptr);
void QVideoWidgetControl_DisconnectBrightnessChanged(QtObjectPtr ptr);
int QVideoWidgetControl_Contrast(QtObjectPtr ptr);
void QVideoWidgetControl_ConnectContrastChanged(QtObjectPtr ptr);
void QVideoWidgetControl_DisconnectContrastChanged(QtObjectPtr ptr);
void QVideoWidgetControl_ConnectFullScreenChanged(QtObjectPtr ptr);
void QVideoWidgetControl_DisconnectFullScreenChanged(QtObjectPtr ptr);
int QVideoWidgetControl_Hue(QtObjectPtr ptr);
void QVideoWidgetControl_ConnectHueChanged(QtObjectPtr ptr);
void QVideoWidgetControl_DisconnectHueChanged(QtObjectPtr ptr);
int QVideoWidgetControl_IsFullScreen(QtObjectPtr ptr);
int QVideoWidgetControl_Saturation(QtObjectPtr ptr);
void QVideoWidgetControl_ConnectSaturationChanged(QtObjectPtr ptr);
void QVideoWidgetControl_DisconnectSaturationChanged(QtObjectPtr ptr);
void QVideoWidgetControl_SetAspectRatioMode(QtObjectPtr ptr, int mode);
void QVideoWidgetControl_SetBrightness(QtObjectPtr ptr, int brightness);
void QVideoWidgetControl_SetContrast(QtObjectPtr ptr, int contrast);
void QVideoWidgetControl_SetFullScreen(QtObjectPtr ptr, int fullScreen);
void QVideoWidgetControl_SetHue(QtObjectPtr ptr, int hue);
void QVideoWidgetControl_SetSaturation(QtObjectPtr ptr, int saturation);
QtObjectPtr QVideoWidgetControl_VideoWidget(QtObjectPtr ptr);
void QVideoWidgetControl_DestroyQVideoWidgetControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif