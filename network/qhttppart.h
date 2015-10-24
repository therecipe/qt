#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QHttpPart_NewQHttpPart();
QtObjectPtr QHttpPart_NewQHttpPart2(QtObjectPtr other);
void QHttpPart_SetBody(QtObjectPtr ptr, QtObjectPtr body);
void QHttpPart_SetBodyDevice(QtObjectPtr ptr, QtObjectPtr device);
void QHttpPart_SetHeader(QtObjectPtr ptr, int header, char* value);
void QHttpPart_SetRawHeader(QtObjectPtr ptr, QtObjectPtr headerName, QtObjectPtr headerValue);
void QHttpPart_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QHttpPart_DestroyQHttpPart(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif