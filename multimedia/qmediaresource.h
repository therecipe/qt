#ifdef __cplusplus
extern "C" {
#endif

void* QMediaResource_NewQMediaResource();
void* QMediaResource_NewQMediaResource4(void* other);
void* QMediaResource_NewQMediaResource3(void* request, char* mimeType);
void* QMediaResource_NewQMediaResource2(void* url, char* mimeType);
int QMediaResource_AudioBitRate(void* ptr);
char* QMediaResource_AudioCodec(void* ptr);
int QMediaResource_ChannelCount(void* ptr);
int QMediaResource_IsNull(void* ptr);
char* QMediaResource_Language(void* ptr);
char* QMediaResource_MimeType(void* ptr);
int QMediaResource_SampleRate(void* ptr);
void QMediaResource_SetAudioBitRate(void* ptr, int rate);
void QMediaResource_SetAudioCodec(void* ptr, char* codec);
void QMediaResource_SetChannelCount(void* ptr, int channels);
void QMediaResource_SetLanguage(void* ptr, char* language);
void QMediaResource_SetResolution(void* ptr, void* resolution);
void QMediaResource_SetResolution2(void* ptr, int width, int height);
void QMediaResource_SetSampleRate(void* ptr, int sampleRate);
void QMediaResource_SetVideoBitRate(void* ptr, int rate);
void QMediaResource_SetVideoCodec(void* ptr, char* codec);
int QMediaResource_VideoBitRate(void* ptr);
char* QMediaResource_VideoCodec(void* ptr);
void QMediaResource_DestroyQMediaResource(void* ptr);

#ifdef __cplusplus
}
#endif