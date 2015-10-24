#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAudioRecorder_NewQAudioRecorder(QtObjectPtr parent);
char* QAudioRecorder_AudioInput(QtObjectPtr ptr);
void QAudioRecorder_ConnectAudioInputChanged(QtObjectPtr ptr);
void QAudioRecorder_DisconnectAudioInputChanged(QtObjectPtr ptr);
char* QAudioRecorder_AudioInputDescription(QtObjectPtr ptr, char* name);
char* QAudioRecorder_AudioInputs(QtObjectPtr ptr);
void QAudioRecorder_ConnectAvailableAudioInputsChanged(QtObjectPtr ptr);
void QAudioRecorder_DisconnectAvailableAudioInputsChanged(QtObjectPtr ptr);
char* QAudioRecorder_DefaultAudioInput(QtObjectPtr ptr);
void QAudioRecorder_SetAudioInput(QtObjectPtr ptr, char* name);
void QAudioRecorder_DestroyQAudioRecorder(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif