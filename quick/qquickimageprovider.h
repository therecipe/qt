#ifdef __cplusplus
extern "C" {
#endif

void* QQuickImageProvider_NewQQuickImageProvider(int ty, int flags);
int QQuickImageProvider_Flags(void* ptr);
int QQuickImageProvider_ImageType(void* ptr);
void QQuickImageProvider_DestroyQQuickImageProvider(void* ptr);

#ifdef __cplusplus
}
#endif