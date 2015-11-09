#ifdef __cplusplus
extern "C" {
#endif

void* QAudioRecorder_NewQAudioRecorder(void* parent);
char* QAudioRecorder_AudioInput(void* ptr);
void QAudioRecorder_ConnectAudioInputChanged(void* ptr);
void QAudioRecorder_DisconnectAudioInputChanged(void* ptr);
char* QAudioRecorder_AudioInputDescription(void* ptr, char* name);
char* QAudioRecorder_AudioInputs(void* ptr);
void QAudioRecorder_ConnectAvailableAudioInputsChanged(void* ptr);
void QAudioRecorder_DisconnectAvailableAudioInputsChanged(void* ptr);
char* QAudioRecorder_DefaultAudioInput(void* ptr);
void QAudioRecorder_SetAudioInput(void* ptr, char* name);
void QAudioRecorder_DestroyQAudioRecorder(void* ptr);

#ifdef __cplusplus
}
#endif