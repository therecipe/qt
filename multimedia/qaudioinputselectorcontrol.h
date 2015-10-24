#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QAudioInputSelectorControl_ActiveInput(QtObjectPtr ptr);
void QAudioInputSelectorControl_ConnectActiveInputChanged(QtObjectPtr ptr);
void QAudioInputSelectorControl_DisconnectActiveInputChanged(QtObjectPtr ptr);
void QAudioInputSelectorControl_ConnectAvailableInputsChanged(QtObjectPtr ptr);
void QAudioInputSelectorControl_DisconnectAvailableInputsChanged(QtObjectPtr ptr);
char* QAudioInputSelectorControl_DefaultInput(QtObjectPtr ptr);
char* QAudioInputSelectorControl_InputDescription(QtObjectPtr ptr, char* name);
void QAudioInputSelectorControl_SetActiveInput(QtObjectPtr ptr, char* name);
void QAudioInputSelectorControl_DestroyQAudioInputSelectorControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif