#ifdef __cplusplus
extern "C" {
#endif

void* QHttpMultiPart_NewQHttpMultiPart2(int contentType, void* parent);
void* QHttpMultiPart_NewQHttpMultiPart(void* parent);
void QHttpMultiPart_Append(void* ptr, void* httpPart);
void* QHttpMultiPart_Boundary(void* ptr);
void QHttpMultiPart_SetBoundary(void* ptr, void* boundary);
void QHttpMultiPart_SetContentType(void* ptr, int contentType);
void QHttpMultiPart_DestroyQHttpMultiPart(void* ptr);

#ifdef __cplusplus
}
#endif