#ifdef __cplusplus
extern "C" {
#endif

void* QMediaContent_NewQMediaContent();
void* QMediaContent_NewQMediaContent7(void* playlist, void* contentUrl, int takeOwnership);
void* QMediaContent_NewQMediaContent6(void* other);
void* QMediaContent_NewQMediaContent4(void* resource);
void* QMediaContent_NewQMediaContent3(void* request);
void* QMediaContent_NewQMediaContent2(void* url);
int QMediaContent_IsNull(void* ptr);
void* QMediaContent_Playlist(void* ptr);
void QMediaContent_DestroyQMediaContent(void* ptr);

#ifdef __cplusplus
}
#endif