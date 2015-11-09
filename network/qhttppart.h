#ifdef __cplusplus
extern "C" {
#endif

void* QHttpPart_NewQHttpPart();
void* QHttpPart_NewQHttpPart2(void* other);
void QHttpPart_SetBody(void* ptr, void* body);
void QHttpPart_SetBodyDevice(void* ptr, void* device);
void QHttpPart_SetHeader(void* ptr, int header, void* value);
void QHttpPart_SetRawHeader(void* ptr, void* headerName, void* headerValue);
void QHttpPart_Swap(void* ptr, void* other);
void QHttpPart_DestroyQHttpPart(void* ptr);

#ifdef __cplusplus
}
#endif