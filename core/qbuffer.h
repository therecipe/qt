#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBuffer_NewQBuffer2(QtObjectPtr byteArray, QtObjectPtr parent);
QtObjectPtr QBuffer_NewQBuffer(QtObjectPtr parent);
int QBuffer_AtEnd(QtObjectPtr ptr);
int QBuffer_CanReadLine(QtObjectPtr ptr);
void QBuffer_Close(QtObjectPtr ptr);
int QBuffer_Open(QtObjectPtr ptr, int flags);
void QBuffer_SetBuffer(QtObjectPtr ptr, QtObjectPtr byteArray);
void QBuffer_SetData(QtObjectPtr ptr, QtObjectPtr data);
void QBuffer_SetData2(QtObjectPtr ptr, char* data, int size);
void QBuffer_DestroyQBuffer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif