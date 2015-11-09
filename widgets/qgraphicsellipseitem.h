#ifdef __cplusplus
extern "C" {
#endif

int QGraphicsEllipseItem_Contains(void* ptr, void* point);
int QGraphicsEllipseItem_IsObscuredBy(void* ptr, void* item);
void QGraphicsEllipseItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsEllipseItem_SetRect(void* ptr, void* rect);
void QGraphicsEllipseItem_SetRect2(void* ptr, double x, double y, double width, double height);
void QGraphicsEllipseItem_SetSpanAngle(void* ptr, int angle);
void QGraphicsEllipseItem_SetStartAngle(void* ptr, int angle);
int QGraphicsEllipseItem_SpanAngle(void* ptr);
int QGraphicsEllipseItem_StartAngle(void* ptr);
int QGraphicsEllipseItem_Type(void* ptr);
void QGraphicsEllipseItem_DestroyQGraphicsEllipseItem(void* ptr);

#ifdef __cplusplus
}
#endif