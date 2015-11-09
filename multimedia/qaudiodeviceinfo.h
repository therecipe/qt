#ifdef __cplusplus
extern "C" {
#endif

void* QAudioDeviceInfo_NewQAudioDeviceInfo();
void* QAudioDeviceInfo_NewQAudioDeviceInfo2(void* other);
char* QAudioDeviceInfo_DeviceName(void* ptr);
int QAudioDeviceInfo_IsFormatSupported(void* ptr, void* settings);
int QAudioDeviceInfo_IsNull(void* ptr);
char* QAudioDeviceInfo_SupportedCodecs(void* ptr);
void QAudioDeviceInfo_DestroyQAudioDeviceInfo(void* ptr);

#ifdef __cplusplus
}
#endif