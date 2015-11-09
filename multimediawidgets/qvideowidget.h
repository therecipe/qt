#ifdef __cplusplus
extern "C" {
#endif

int QVideoWidget_AspectRatioMode(void* ptr);
int QVideoWidget_Brightness(void* ptr);
int QVideoWidget_Contrast(void* ptr);
int QVideoWidget_Hue(void* ptr);
void* QVideoWidget_MediaObject(void* ptr);
int QVideoWidget_Saturation(void* ptr);
void QVideoWidget_SetAspectRatioMode(void* ptr, int mode);
void QVideoWidget_SetBrightness(void* ptr, int brightness);
void QVideoWidget_SetContrast(void* ptr, int contrast);
void QVideoWidget_SetFullScreen(void* ptr, int fullScreen);
void QVideoWidget_SetHue(void* ptr, int hue);
void QVideoWidget_SetSaturation(void* ptr, int saturation);
void* QVideoWidget_NewQVideoWidget(void* parent);
void QVideoWidget_ConnectBrightnessChanged(void* ptr);
void QVideoWidget_DisconnectBrightnessChanged(void* ptr);
void QVideoWidget_ConnectContrastChanged(void* ptr);
void QVideoWidget_DisconnectContrastChanged(void* ptr);
void QVideoWidget_ConnectFullScreenChanged(void* ptr);
void QVideoWidget_DisconnectFullScreenChanged(void* ptr);
void QVideoWidget_ConnectHueChanged(void* ptr);
void QVideoWidget_DisconnectHueChanged(void* ptr);
int QVideoWidget_IsFullScreen(void* ptr);
void QVideoWidget_ConnectSaturationChanged(void* ptr);
void QVideoWidget_DisconnectSaturationChanged(void* ptr);
void QVideoWidget_DestroyQVideoWidget(void* ptr);

#ifdef __cplusplus
}
#endif