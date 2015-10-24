#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QCameraFocusZone_NewQCameraFocusZone(QtObjectPtr other);
int QCameraFocusZone_IsValid(QtObjectPtr ptr);
int QCameraFocusZone_Status(QtObjectPtr ptr);
void QCameraFocusZone_DestroyQCameraFocusZone(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif