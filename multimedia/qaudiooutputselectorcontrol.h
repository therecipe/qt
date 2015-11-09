#ifdef __cplusplus
extern "C" {
#endif

char* QAudioOutputSelectorControl_ActiveOutput(void* ptr);
void QAudioOutputSelectorControl_ConnectActiveOutputChanged(void* ptr);
void QAudioOutputSelectorControl_DisconnectActiveOutputChanged(void* ptr);
void QAudioOutputSelectorControl_ConnectAvailableOutputsChanged(void* ptr);
void QAudioOutputSelectorControl_DisconnectAvailableOutputsChanged(void* ptr);
char* QAudioOutputSelectorControl_DefaultOutput(void* ptr);
char* QAudioOutputSelectorControl_OutputDescription(void* ptr, char* name);
void QAudioOutputSelectorControl_SetActiveOutput(void* ptr, char* name);
void QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(void* ptr);

#ifdef __cplusplus
}
#endif