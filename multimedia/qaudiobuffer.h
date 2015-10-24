#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAudioBuffer_NewQAudioBuffer();
QtObjectPtr QAudioBuffer_NewQAudioBuffer3(QtObjectPtr other);
int QAudioBuffer_ByteCount(QtObjectPtr ptr);
void QAudioBuffer_ConstData(QtObjectPtr ptr);
void QAudioBuffer_Data2(QtObjectPtr ptr);
void QAudioBuffer_Data(QtObjectPtr ptr);
int QAudioBuffer_FrameCount(QtObjectPtr ptr);
int QAudioBuffer_IsValid(QtObjectPtr ptr);
int QAudioBuffer_SampleCount(QtObjectPtr ptr);
void QAudioBuffer_DestroyQAudioBuffer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif