#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QHttpMultiPart_NewQHttpMultiPart2(int contentType, QtObjectPtr parent);
QtObjectPtr QHttpMultiPart_NewQHttpMultiPart(QtObjectPtr parent);
void QHttpMultiPart_Append(QtObjectPtr ptr, QtObjectPtr httpPart);
void QHttpMultiPart_SetBoundary(QtObjectPtr ptr, QtObjectPtr boundary);
void QHttpMultiPart_SetContentType(QtObjectPtr ptr, int contentType);
void QHttpMultiPart_DestroyQHttpMultiPart(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif