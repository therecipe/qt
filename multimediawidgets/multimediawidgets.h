#ifdef __cplusplus
extern "C" {
#endif

void* QCameraViewfinder_NewQCameraViewfinder(void* parent);
void* QCameraViewfinder_MediaObject(void* ptr);
void QCameraViewfinder_DestroyQCameraViewfinder(void* ptr);
void* QGraphicsVideoItem_NewQGraphicsVideoItem(void* parent);
int QGraphicsVideoItem_AspectRatioMode(void* ptr);
void* QGraphicsVideoItem_MediaObject(void* ptr);
void QGraphicsVideoItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsVideoItem_SetAspectRatioMode(void* ptr, int mode);
void QGraphicsVideoItem_SetOffset(void* ptr, void* offset);
void QGraphicsVideoItem_SetSize(void* ptr, void* size);
void QGraphicsVideoItem_DestroyQGraphicsVideoItem(void* ptr);
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