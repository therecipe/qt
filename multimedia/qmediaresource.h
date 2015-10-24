#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMediaResource_NewQMediaResource();
QtObjectPtr QMediaResource_NewQMediaResource4(QtObjectPtr other);
QtObjectPtr QMediaResource_NewQMediaResource3(QtObjectPtr request, char* mimeType);
QtObjectPtr QMediaResource_NewQMediaResource2(char* url, char* mimeType);
int QMediaResource_AudioBitRate(QtObjectPtr ptr);
char* QMediaResource_AudioCodec(QtObjectPtr ptr);
int QMediaResource_ChannelCount(QtObjectPtr ptr);
int QMediaResource_IsNull(QtObjectPtr ptr);
char* QMediaResource_Language(QtObjectPtr ptr);
char* QMediaResource_MimeType(QtObjectPtr ptr);
int QMediaResource_SampleRate(QtObjectPtr ptr);
void QMediaResource_SetAudioBitRate(QtObjectPtr ptr, int rate);
void QMediaResource_SetAudioCodec(QtObjectPtr ptr, char* codec);
void QMediaResource_SetChannelCount(QtObjectPtr ptr, int channels);
void QMediaResource_SetLanguage(QtObjectPtr ptr, char* language);
void QMediaResource_SetResolution(QtObjectPtr ptr, QtObjectPtr resolution);
void QMediaResource_SetResolution2(QtObjectPtr ptr, int width, int height);
void QMediaResource_SetSampleRate(QtObjectPtr ptr, int sampleRate);
void QMediaResource_SetVideoBitRate(QtObjectPtr ptr, int rate);
void QMediaResource_SetVideoCodec(QtObjectPtr ptr, char* codec);
char* QMediaResource_Url(QtObjectPtr ptr);
int QMediaResource_VideoBitRate(QtObjectPtr ptr);
char* QMediaResource_VideoCodec(QtObjectPtr ptr);
void QMediaResource_DestroyQMediaResource(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif