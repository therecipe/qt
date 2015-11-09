#ifdef __cplusplus
extern "C" {
#endif

void* QAudioFormat_NewQAudioFormat();
void* QAudioFormat_NewQAudioFormat2(void* other);
int QAudioFormat_ByteOrder(void* ptr);
int QAudioFormat_BytesPerFrame(void* ptr);
int QAudioFormat_ChannelCount(void* ptr);
char* QAudioFormat_Codec(void* ptr);
int QAudioFormat_IsValid(void* ptr);
int QAudioFormat_SampleRate(void* ptr);
int QAudioFormat_SampleSize(void* ptr);
int QAudioFormat_SampleType(void* ptr);
void QAudioFormat_SetByteOrder(void* ptr, int byteOrder);
void QAudioFormat_SetChannelCount(void* ptr, int channels);
void QAudioFormat_SetCodec(void* ptr, char* codec);
void QAudioFormat_SetSampleRate(void* ptr, int samplerate);
void QAudioFormat_SetSampleSize(void* ptr, int sampleSize);
void QAudioFormat_SetSampleType(void* ptr, int sampleType);
void QAudioFormat_DestroyQAudioFormat(void* ptr);

#ifdef __cplusplus
}
#endif