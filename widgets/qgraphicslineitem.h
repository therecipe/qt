#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsLineItem_NewQGraphicsLineItem(void* parent);
void* QGraphicsLineItem_NewQGraphicsLineItem2(void* line, void* parent);
void* QGraphicsLineItem_NewQGraphicsLineItem3(double x1, double y1, double x2, double y2, void* parent);
int QGraphicsLineItem_Contains(void* ptr, void* point);
int QGraphicsLineItem_IsObscuredBy(void* ptr, void* item);
void QGraphicsLineItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsLineItem_SetLine(void* ptr, void* line);
void QGraphicsLineItem_SetLine2(void* ptr, double x1, double y1, double x2, double y2);
void QGraphicsLineItem_SetPen(void* ptr, void* pen);
int QGraphicsLineItem_Type(void* ptr);
void QGraphicsLineItem_DestroyQGraphicsLineItem(void* ptr);

#ifdef __cplusplus
}
#endif