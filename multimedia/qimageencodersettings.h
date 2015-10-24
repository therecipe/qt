#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QImageEncoderSettings_NewQImageEncoderSettings();
QtObjectPtr QImageEncoderSettings_NewQImageEncoderSettings2(QtObjectPtr other);
char* QImageEncoderSettings_Codec(QtObjectPtr ptr);
char* QImageEncoderSettings_EncodingOption(QtObjectPtr ptr, char* option);
int QImageEncoderSettings_IsNull(QtObjectPtr ptr);
void QImageEncoderSettings_SetCodec(QtObjectPtr ptr, char* codec);
void QImageEncoderSettings_SetEncodingOption(QtObjectPtr ptr, char* option, char* value);
void QImageEncoderSettings_SetResolution(QtObjectPtr ptr, QtObjectPtr resolution);
void QImageEncoderSettings_SetResolution2(QtObjectPtr ptr, int width, int height);
void QImageEncoderSettings_DestroyQImageEncoderSettings(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif