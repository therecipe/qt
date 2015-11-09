#ifdef __cplusplus
extern "C" {
#endif

void QGraphicsRectItem_SetRect(void* ptr, void* rectangle);
int QGraphicsRectItem_Contains(void* ptr, void* point);
int QGraphicsRectItem_IsObscuredBy(void* ptr, void* item);
void QGraphicsRectItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsRectItem_SetRect2(void* ptr, double x, double y, double width, double height);
int QGraphicsRectItem_Type(void* ptr);
void QGraphicsRectItem_DestroyQGraphicsRectItem(void* ptr);

#ifdef __cplusplus
}
#endif