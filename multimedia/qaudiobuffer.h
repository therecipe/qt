#ifdef __cplusplus
extern "C" {
#endif

void* QAudioBuffer_NewQAudioBuffer();
void* QAudioBuffer_NewQAudioBuffer3(void* other);
int QAudioBuffer_ByteCount(void* ptr);
void* QAudioBuffer_ConstData(void* ptr);
void* QAudioBuffer_Data2(void* ptr);
void* QAudioBuffer_Data(void* ptr);
int QAudioBuffer_FrameCount(void* ptr);
int QAudioBuffer_IsValid(void* ptr);
int QAudioBuffer_SampleCount(void* ptr);
void QAudioBuffer_DestroyQAudioBuffer(void* ptr);

#ifdef __cplusplus
}
#endif