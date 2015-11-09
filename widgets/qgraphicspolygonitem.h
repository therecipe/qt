#ifdef __cplusplus
extern "C" {
#endif

int QGraphicsPolygonItem_Contains(void* ptr, void* point);
int QGraphicsPolygonItem_FillRule(void* ptr);
int QGraphicsPolygonItem_IsObscuredBy(void* ptr, void* item);
void QGraphicsPolygonItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsPolygonItem_SetFillRule(void* ptr, int rule);
void QGraphicsPolygonItem_SetPolygon(void* ptr, void* polygon);
int QGraphicsPolygonItem_Type(void* ptr);
void QGraphicsPolygonItem_DestroyQGraphicsPolygonItem(void* ptr);

#ifdef __cplusplus
}
#endif