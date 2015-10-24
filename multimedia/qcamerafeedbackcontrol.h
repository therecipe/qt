#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraFeedbackControl_IsEventFeedbackEnabled(QtObjectPtr ptr, int event);
int QCameraFeedbackControl_IsEventFeedbackLocked(QtObjectPtr ptr, int event);
void QCameraFeedbackControl_ResetEventFeedback(QtObjectPtr ptr, int event);
int QCameraFeedbackControl_SetEventFeedbackEnabled(QtObjectPtr ptr, int event, int enabled);
int QCameraFeedbackControl_SetEventFeedbackSound(QtObjectPtr ptr, int event, char* filePath);
void QCameraFeedbackControl_DestroyQCameraFeedbackControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif