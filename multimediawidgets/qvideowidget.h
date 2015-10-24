#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QVideoWidget_AspectRatioMode(QtObjectPtr ptr);
int QVideoWidget_Brightness(QtObjectPtr ptr);
int QVideoWidget_Contrast(QtObjectPtr ptr);
int QVideoWidget_Hue(QtObjectPtr ptr);
QtObjectPtr QVideoWidget_MediaObject(QtObjectPtr ptr);
int QVideoWidget_Saturation(QtObjectPtr ptr);
void QVideoWidget_SetAspectRatioMode(QtObjectPtr ptr, int mode);
void QVideoWidget_SetBrightness(QtObjectPtr ptr, int brightness);
void QVideoWidget_SetContrast(QtObjectPtr ptr, int contrast);
void QVideoWidget_SetFullScreen(QtObjectPtr ptr, int fullScreen);
void QVideoWidget_SetHue(QtObjectPtr ptr, int hue);
void QVideoWidget_SetSaturation(QtObjectPtr ptr, int saturation);
QtObjectPtr QVideoWidget_NewQVideoWidget(QtObjectPtr parent);
void QVideoWidget_ConnectBrightnessChanged(QtObjectPtr ptr);
void QVideoWidget_DisconnectBrightnessChanged(QtObjectPtr ptr);
void QVideoWidget_ConnectContrastChanged(QtObjectPtr ptr);
void QVideoWidget_DisconnectContrastChanged(QtObjectPtr ptr);
void QVideoWidget_ConnectFullScreenChanged(QtObjectPtr ptr);
void QVideoWidget_DisconnectFullScreenChanged(QtObjectPtr ptr);
void QVideoWidget_ConnectHueChanged(QtObjectPtr ptr);
void QVideoWidget_DisconnectHueChanged(QtObjectPtr ptr);
int QVideoWidget_IsFullScreen(QtObjectPtr ptr);
void QVideoWidget_ConnectSaturationChanged(QtObjectPtr ptr);
void QVideoWidget_DisconnectSaturationChanged(QtObjectPtr ptr);
void QVideoWidget_DestroyQVideoWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif