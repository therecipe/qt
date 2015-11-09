#ifdef __cplusplus
extern "C" {
#endif

void* QImageEncoderSettings_NewQImageEncoderSettings();
void* QImageEncoderSettings_NewQImageEncoderSettings2(void* other);
char* QImageEncoderSettings_Codec(void* ptr);
void* QImageEncoderSettings_EncodingOption(void* ptr, char* option);
int QImageEncoderSettings_IsNull(void* ptr);
void QImageEncoderSettings_SetCodec(void* ptr, char* codec);
void QImageEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value);
void QImageEncoderSettings_SetResolution(void* ptr, void* resolution);
void QImageEncoderSettings_SetResolution2(void* ptr, int width, int height);
void QImageEncoderSettings_DestroyQImageEncoderSettings(void* ptr);

#ifdef __cplusplus
}
#endif