#ifdef __cplusplus
extern "C" {
#endif

void* QAudioEncoderSettings_NewQAudioEncoderSettings();
void* QAudioEncoderSettings_NewQAudioEncoderSettings2(void* other);
int QAudioEncoderSettings_BitRate(void* ptr);
int QAudioEncoderSettings_ChannelCount(void* ptr);
char* QAudioEncoderSettings_Codec(void* ptr);
void* QAudioEncoderSettings_EncodingOption(void* ptr, char* option);
int QAudioEncoderSettings_IsNull(void* ptr);
int QAudioEncoderSettings_SampleRate(void* ptr);
void QAudioEncoderSettings_SetBitRate(void* ptr, int rate);
void QAudioEncoderSettings_SetChannelCount(void* ptr, int channels);
void QAudioEncoderSettings_SetCodec(void* ptr, char* codec);
void QAudioEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value);
void QAudioEncoderSettings_SetSampleRate(void* ptr, int rate);
void QAudioEncoderSettings_DestroyQAudioEncoderSettings(void* ptr);

#ifdef __cplusplus
}
#endif