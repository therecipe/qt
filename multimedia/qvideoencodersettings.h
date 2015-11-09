#ifdef __cplusplus
extern "C" {
#endif

void QVideoEncoderSettings_SetFrameRate(void* ptr, double rate);
void* QVideoEncoderSettings_NewQVideoEncoderSettings();
void* QVideoEncoderSettings_NewQVideoEncoderSettings2(void* other);
int QVideoEncoderSettings_BitRate(void* ptr);
char* QVideoEncoderSettings_Codec(void* ptr);
void* QVideoEncoderSettings_EncodingOption(void* ptr, char* option);
double QVideoEncoderSettings_FrameRate(void* ptr);
int QVideoEncoderSettings_IsNull(void* ptr);
void QVideoEncoderSettings_SetBitRate(void* ptr, int value);
void QVideoEncoderSettings_SetCodec(void* ptr, char* codec);
void QVideoEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value);
void QVideoEncoderSettings_SetResolution(void* ptr, void* resolution);
void QVideoEncoderSettings_SetResolution2(void* ptr, int width, int height);
void QVideoEncoderSettings_DestroyQVideoEncoderSettings(void* ptr);

#ifdef __cplusplus
}
#endif