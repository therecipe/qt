#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAudioFormat_NewQAudioFormat();
QtObjectPtr QAudioFormat_NewQAudioFormat2(QtObjectPtr other);
int QAudioFormat_ByteOrder(QtObjectPtr ptr);
int QAudioFormat_BytesPerFrame(QtObjectPtr ptr);
int QAudioFormat_ChannelCount(QtObjectPtr ptr);
char* QAudioFormat_Codec(QtObjectPtr ptr);
int QAudioFormat_IsValid(QtObjectPtr ptr);
int QAudioFormat_SampleRate(QtObjectPtr ptr);
int QAudioFormat_SampleSize(QtObjectPtr ptr);
int QAudioFormat_SampleType(QtObjectPtr ptr);
void QAudioFormat_SetByteOrder(QtObjectPtr ptr, int byteOrder);
void QAudioFormat_SetChannelCount(QtObjectPtr ptr, int channels);
void QAudioFormat_SetCodec(QtObjectPtr ptr, char* codec);
void QAudioFormat_SetSampleRate(QtObjectPtr ptr, int samplerate);
void QAudioFormat_SetSampleSize(QtObjectPtr ptr, int sampleSize);
void QAudioFormat_SetSampleType(QtObjectPtr ptr, int sampleType);
void QAudioFormat_DestroyQAudioFormat(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif