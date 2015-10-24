#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QVideoEncoderSettings_NewQVideoEncoderSettings();
QtObjectPtr QVideoEncoderSettings_NewQVideoEncoderSettings2(QtObjectPtr other);
int QVideoEncoderSettings_BitRate(QtObjectPtr ptr);
char* QVideoEncoderSettings_Codec(QtObjectPtr ptr);
char* QVideoEncoderSettings_EncodingOption(QtObjectPtr ptr, char* option);
int QVideoEncoderSettings_IsNull(QtObjectPtr ptr);
void QVideoEncoderSettings_SetBitRate(QtObjectPtr ptr, int value);
void QVideoEncoderSettings_SetCodec(QtObjectPtr ptr, char* codec);
void QVideoEncoderSettings_SetEncodingOption(QtObjectPtr ptr, char* option, char* value);
void QVideoEncoderSettings_SetResolution(QtObjectPtr ptr, QtObjectPtr resolution);
void QVideoEncoderSettings_SetResolution2(QtObjectPtr ptr, int width, int height);
void QVideoEncoderSettings_DestroyQVideoEncoderSettings(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif