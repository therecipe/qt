#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QAudioOutputSelectorControl_ActiveOutput(QtObjectPtr ptr);
void QAudioOutputSelectorControl_ConnectActiveOutputChanged(QtObjectPtr ptr);
void QAudioOutputSelectorControl_DisconnectActiveOutputChanged(QtObjectPtr ptr);
void QAudioOutputSelectorControl_ConnectAvailableOutputsChanged(QtObjectPtr ptr);
void QAudioOutputSelectorControl_DisconnectAvailableOutputsChanged(QtObjectPtr ptr);
char* QAudioOutputSelectorControl_DefaultOutput(QtObjectPtr ptr);
char* QAudioOutputSelectorControl_OutputDescription(QtObjectPtr ptr, char* name);
void QAudioOutputSelectorControl_SetActiveOutput(QtObjectPtr ptr, char* name);
void QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif