#ifdef __cplusplus
extern "C" {
#endif

char* QImageEncoderControl_ImageCodecDescription(void* ptr, char* codec);
void QImageEncoderControl_SetImageSettings(void* ptr, void* settings);
char* QImageEncoderControl_SupportedImageCodecs(void* ptr);
void QImageEncoderControl_DestroyQImageEncoderControl(void* ptr);

#ifdef __cplusplus
}
#endif