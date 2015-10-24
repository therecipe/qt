#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAudioDeviceInfo_NewQAudioDeviceInfo();
QtObjectPtr QAudioDeviceInfo_NewQAudioDeviceInfo2(QtObjectPtr other);
char* QAudioDeviceInfo_DeviceName(QtObjectPtr ptr);
int QAudioDeviceInfo_IsFormatSupported(QtObjectPtr ptr, QtObjectPtr settings);
int QAudioDeviceInfo_IsNull(QtObjectPtr ptr);
char* QAudioDeviceInfo_SupportedCodecs(QtObjectPtr ptr);
void QAudioDeviceInfo_DestroyQAudioDeviceInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif