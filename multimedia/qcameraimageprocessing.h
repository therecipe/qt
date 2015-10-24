#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCameraImageProcessing_ColorFilter(QtObjectPtr ptr);
int QCameraImageProcessing_IsAvailable(QtObjectPtr ptr);
int QCameraImageProcessing_IsColorFilterSupported(QtObjectPtr ptr, int filter);
int QCameraImageProcessing_IsWhiteBalanceModeSupported(QtObjectPtr ptr, int mode);
void QCameraImageProcessing_SetColorFilter(QtObjectPtr ptr, int filter);
void QCameraImageProcessing_SetWhiteBalanceMode(QtObjectPtr ptr, int mode);
int QCameraImageProcessing_WhiteBalanceMode(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif