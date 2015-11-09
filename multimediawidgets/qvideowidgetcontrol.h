#ifdef __cplusplus
extern "C" {
#endif

int QVideoWidgetControl_AspectRatioMode(void* ptr);
int QVideoWidgetControl_Brightness(void* ptr);
void QVideoWidgetControl_ConnectBrightnessChanged(void* ptr);
void QVideoWidgetControl_DisconnectBrightnessChanged(void* ptr);
int QVideoWidgetControl_Contrast(void* ptr);
void QVideoWidgetControl_ConnectContrastChanged(void* ptr);
void QVideoWidgetControl_DisconnectContrastChanged(void* ptr);
void QVideoWidgetControl_ConnectFullScreenChanged(void* ptr);
void QVideoWidgetControl_DisconnectFullScreenChanged(void* ptr);
int QVideoWidgetControl_Hue(void* ptr);
void QVideoWidgetControl_ConnectHueChanged(void* ptr);
void QVideoWidgetControl_DisconnectHueChanged(void* ptr);
int QVideoWidgetControl_IsFullScreen(void* ptr);
int QVideoWidgetControl_Saturation(void* ptr);
void QVideoWidgetControl_ConnectSaturationChanged(void* ptr);
void QVideoWidgetControl_DisconnectSaturationChanged(void* ptr);
void QVideoWidgetControl_SetAspectRatioMode(void* ptr, int mode);
void QVideoWidgetControl_SetBrightness(void* ptr, int brightness);
void QVideoWidgetControl_SetContrast(void* ptr, int contrast);
void QVideoWidgetControl_SetFullScreen(void* ptr, int fullScreen);
void QVideoWidgetControl_SetHue(void* ptr, int hue);
void QVideoWidgetControl_SetSaturation(void* ptr, int saturation);
void* QVideoWidgetControl_VideoWidget(void* ptr);
void QVideoWidgetControl_DestroyQVideoWidgetControl(void* ptr);

#ifdef __cplusplus
}
#endif