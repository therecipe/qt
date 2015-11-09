#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsPixmapItem_NewQGraphicsPixmapItem(void* parent);
void* QGraphicsPixmapItem_NewQGraphicsPixmapItem2(void* pixmap, void* parent);
int QGraphicsPixmapItem_Contains(void* ptr, void* point);
int QGraphicsPixmapItem_IsObscuredBy(void* ptr, void* item);
void QGraphicsPixmapItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsPixmapItem_SetOffset(void* ptr, void* offset);
void QGraphicsPixmapItem_SetOffset2(void* ptr, double x, double y);
void QGraphicsPixmapItem_SetPixmap(void* ptr, void* pixmap);
void QGraphicsPixmapItem_SetShapeMode(void* ptr, int mode);
void QGraphicsPixmapItem_SetTransformationMode(void* ptr, int mode);
int QGraphicsPixmapItem_ShapeMode(void* ptr);
int QGraphicsPixmapItem_TransformationMode(void* ptr);
int QGraphicsPixmapItem_Type(void* ptr);
void QGraphicsPixmapItem_DestroyQGraphicsPixmapItem(void* ptr);

#ifdef __cplusplus
}
#endif