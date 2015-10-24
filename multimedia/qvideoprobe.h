#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QVideoProbe_NewQVideoProbe(QtObjectPtr parent);
void QVideoProbe_ConnectFlush(QtObjectPtr ptr);
void QVideoProbe_DisconnectFlush(QtObjectPtr ptr);
int QVideoProbe_IsActive(QtObjectPtr ptr);
int QVideoProbe_SetSource(QtObjectPtr ptr, QtObjectPtr source);
int QVideoProbe_SetSource2(QtObjectPtr ptr, QtObjectPtr mediaRecorder);
void QVideoProbe_DestroyQVideoProbe(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif