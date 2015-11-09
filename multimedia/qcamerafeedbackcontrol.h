#ifdef __cplusplus
extern "C" {
#endif

int QCameraFeedbackControl_IsEventFeedbackEnabled(void* ptr, int event);
int QCameraFeedbackControl_IsEventFeedbackLocked(void* ptr, int event);
void QCameraFeedbackControl_ResetEventFeedback(void* ptr, int event);
int QCameraFeedbackControl_SetEventFeedbackEnabled(void* ptr, int event, int enabled);
int QCameraFeedbackControl_SetEventFeedbackSound(void* ptr, int event, char* filePath);
void QCameraFeedbackControl_DestroyQCameraFeedbackControl(void* ptr);

#ifdef __cplusplus
}
#endif