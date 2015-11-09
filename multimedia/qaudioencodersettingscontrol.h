#ifdef __cplusplus
extern "C" {
#endif

char* QAudioEncoderSettingsControl_CodecDescription(void* ptr, char* codec);
void QAudioEncoderSettingsControl_SetAudioSettings(void* ptr, void* settings);
char* QAudioEncoderSettingsControl_SupportedAudioCodecs(void* ptr);
void QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(void* ptr);

#ifdef __cplusplus
}
#endif