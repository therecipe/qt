#ifdef __cplusplus
extern "C" {
#endif

void QVideoEncoderSettingsControl_SetVideoSettings(void* ptr, void* settings);
char* QVideoEncoderSettingsControl_SupportedVideoCodecs(void* ptr);
char* QVideoEncoderSettingsControl_VideoCodecDescription(void* ptr, char* codec);
void QVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControl(void* ptr);

#ifdef __cplusplus
}
#endif