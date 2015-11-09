#ifdef __cplusplus
extern "C" {
#endif

void QMediaService_ReleaseControl(void* ptr, void* control);
void* QMediaService_RequestControl(void* ptr, char* interfa);
void* QMediaService_RequestControl2(void* ptr);
void QMediaService_DestroyQMediaService(void* ptr);

#ifdef __cplusplus
}
#endif