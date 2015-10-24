#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QImageEncoderControl_ImageCodecDescription(QtObjectPtr ptr, char* codec);
void QImageEncoderControl_SetImageSettings(QtObjectPtr ptr, QtObjectPtr settings);
char* QImageEncoderControl_SupportedImageCodecs(QtObjectPtr ptr);
void QImageEncoderControl_DestroyQImageEncoderControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif