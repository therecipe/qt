#ifdef __cplusplus
extern "C" {
#endif

int QCameraImageProcessing_ColorFilter(void* ptr);
double QCameraImageProcessing_Contrast(void* ptr);
double QCameraImageProcessing_DenoisingLevel(void* ptr);
int QCameraImageProcessing_IsAvailable(void* ptr);
int QCameraImageProcessing_IsColorFilterSupported(void* ptr, int filter);
int QCameraImageProcessing_IsWhiteBalanceModeSupported(void* ptr, int mode);
double QCameraImageProcessing_ManualWhiteBalance(void* ptr);
double QCameraImageProcessing_Saturation(void* ptr);
void QCameraImageProcessing_SetColorFilter(void* ptr, int filter);
void QCameraImageProcessing_SetContrast(void* ptr, double value);
void QCameraImageProcessing_SetDenoisingLevel(void* ptr, double level);
void QCameraImageProcessing_SetManualWhiteBalance(void* ptr, double colorTemperature);
void QCameraImageProcessing_SetSaturation(void* ptr, double value);
void QCameraImageProcessing_SetSharpeningLevel(void* ptr, double level);
void QCameraImageProcessing_SetWhiteBalanceMode(void* ptr, int mode);
double QCameraImageProcessing_SharpeningLevel(void* ptr);
int QCameraImageProcessing_WhiteBalanceMode(void* ptr);

#ifdef __cplusplus
}
#endif