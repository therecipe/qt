#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QAudioEncoderSettingsControl_CodecDescription(QtObjectPtr ptr, char* codec);
void QAudioEncoderSettingsControl_SetAudioSettings(QtObjectPtr ptr, QtObjectPtr settings);
char* QAudioEncoderSettingsControl_SupportedAudioCodecs(QtObjectPtr ptr);
void QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif