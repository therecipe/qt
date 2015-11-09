#ifdef __cplusplus
extern "C" {
#endif

int QCameraCaptureBufferFormatControl_BufferFormat(void* ptr);
void QCameraCaptureBufferFormatControl_ConnectBufferFormatChanged(void* ptr);
void QCameraCaptureBufferFormatControl_DisconnectBufferFormatChanged(void* ptr);
void QCameraCaptureBufferFormatControl_SetBufferFormat(void* ptr, int format);
void QCameraCaptureBufferFormatControl_DestroyQCameraCaptureBufferFormatControl(void* ptr);

#ifdef __cplusplus
}
#endif