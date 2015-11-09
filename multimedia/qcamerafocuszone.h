#ifdef __cplusplus
extern "C" {
#endif

void* QCameraFocusZone_NewQCameraFocusZone(void* other);
int QCameraFocusZone_IsValid(void* ptr);
int QCameraFocusZone_Status(void* ptr);
void QCameraFocusZone_DestroyQCameraFocusZone(void* ptr);

#ifdef __cplusplus
}
#endif