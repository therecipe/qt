#ifdef __cplusplus
extern "C" {
#endif

char* QAudioInputSelectorControl_ActiveInput(void* ptr);
void QAudioInputSelectorControl_ConnectActiveInputChanged(void* ptr);
void QAudioInputSelectorControl_DisconnectActiveInputChanged(void* ptr);
void QAudioInputSelectorControl_ConnectAvailableInputsChanged(void* ptr);
void QAudioInputSelectorControl_DisconnectAvailableInputsChanged(void* ptr);
char* QAudioInputSelectorControl_DefaultInput(void* ptr);
char* QAudioInputSelectorControl_InputDescription(void* ptr, char* name);
void QAudioInputSelectorControl_SetActiveInput(void* ptr, char* name);
void QAudioInputSelectorControl_DestroyQAudioInputSelectorControl(void* ptr);

#ifdef __cplusplus
}
#endif