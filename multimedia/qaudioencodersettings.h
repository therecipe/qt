#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAudioEncoderSettings_NewQAudioEncoderSettings();
QtObjectPtr QAudioEncoderSettings_NewQAudioEncoderSettings2(QtObjectPtr other);
int QAudioEncoderSettings_BitRate(QtObjectPtr ptr);
int QAudioEncoderSettings_ChannelCount(QtObjectPtr ptr);
char* QAudioEncoderSettings_Codec(QtObjectPtr ptr);
char* QAudioEncoderSettings_EncodingOption(QtObjectPtr ptr, char* option);
int QAudioEncoderSettings_IsNull(QtObjectPtr ptr);
int QAudioEncoderSettings_SampleRate(QtObjectPtr ptr);
void QAudioEncoderSettings_SetBitRate(QtObjectPtr ptr, int rate);
void QAudioEncoderSettings_SetChannelCount(QtObjectPtr ptr, int channels);
void QAudioEncoderSettings_SetCodec(QtObjectPtr ptr, char* codec);
void QAudioEncoderSettings_SetEncodingOption(QtObjectPtr ptr, char* option, char* value);
void QAudioEncoderSettings_SetSampleRate(QtObjectPtr ptr, int rate);
void QAudioEncoderSettings_DestroyQAudioEncoderSettings(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif