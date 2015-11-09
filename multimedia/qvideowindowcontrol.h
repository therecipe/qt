#ifdef __cplusplus
extern "C" {
#endif

int QVideoWindowControl_AspectRatioMode(void* ptr);
int QVideoWindowControl_Brightness(void* ptr);
void QVideoWindowControl_ConnectBrightnessChanged(void* ptr);
void QVideoWindowControl_DisconnectBrightnessChanged(void* ptr);
int QVideoWindowControl_Contrast(void* ptr);
void QVideoWindowControl_ConnectContrastChanged(void* ptr);
void QVideoWindowControl_DisconnectContrastChanged(void* ptr);
void QVideoWindowControl_ConnectFullScreenChanged(void* ptr);
void QVideoWindowControl_DisconnectFullScreenChanged(void* ptr);
int QVideoWindowControl_Hue(void* ptr);
void QVideoWindowControl_ConnectHueChanged(void* ptr);
void QVideoWindowControl_DisconnectHueChanged(void* ptr);
int QVideoWindowControl_IsFullScreen(void* ptr);
void QVideoWindowControl_ConnectNativeSizeChanged(void* ptr);
void QVideoWindowControl_DisconnectNativeSizeChanged(void* ptr);
void QVideoWindowControl_Repaint(void* ptr);
int QVideoWindowControl_Saturation(void* ptr);
void QVideoWindowControl_ConnectSaturationChanged(void* ptr);
void QVideoWindowControl_DisconnectSaturationChanged(void* ptr);
void QVideoWindowControl_SetAspectRatioMode(void* ptr, int mode);
void QVideoWindowControl_SetBrightness(void* ptr, int brightness);
void QVideoWindowControl_SetContrast(void* ptr, int contrast);
void QVideoWindowControl_SetDisplayRect(void* ptr, void* rect);
void QVideoWindowControl_SetFullScreen(void* ptr, int fullScreen);
void QVideoWindowControl_SetHue(void* ptr, int hue);
void QVideoWindowControl_SetSaturation(void* ptr, int saturation);
void QVideoWindowControl_DestroyQVideoWindowControl(void* ptr);

#ifdef __cplusplus
}
#endif