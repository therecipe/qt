#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QVideoEncoderSettingsControl_SetVideoSettings(QtObjectPtr ptr, QtObjectPtr settings);
char* QVideoEncoderSettingsControl_SupportedVideoCodecs(QtObjectPtr ptr);
char* QVideoEncoderSettingsControl_VideoCodecDescription(QtObjectPtr ptr, char* codec);
void QVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif