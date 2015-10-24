#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAudioProbe_NewQAudioProbe(QtObjectPtr parent);
void QAudioProbe_ConnectFlush(QtObjectPtr ptr);
void QAudioProbe_DisconnectFlush(QtObjectPtr ptr);
int QAudioProbe_IsActive(QtObjectPtr ptr);
int QAudioProbe_SetSource(QtObjectPtr ptr, QtObjectPtr source);
int QAudioProbe_SetSource2(QtObjectPtr ptr, QtObjectPtr mediaRecorder);
void QAudioProbe_DestroyQAudioProbe(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif