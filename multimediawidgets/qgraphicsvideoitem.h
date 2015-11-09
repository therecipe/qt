#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsVideoItem_NewQGraphicsVideoItem(void* parent);
int QGraphicsVideoItem_AspectRatioMode(void* ptr);
void* QGraphicsVideoItem_MediaObject(void* ptr);
void QGraphicsVideoItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsVideoItem_SetAspectRatioMode(void* ptr, int mode);
void QGraphicsVideoItem_SetOffset(void* ptr, void* offset);
void QGraphicsVideoItem_SetSize(void* ptr, void* size);
void QGraphicsVideoItem_DestroyQGraphicsVideoItem(void* ptr);

#ifdef __cplusplus
}
#endif