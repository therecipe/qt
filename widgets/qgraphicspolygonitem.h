#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGraphicsPolygonItem_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QGraphicsPolygonItem_FillRule(QtObjectPtr ptr);
int QGraphicsPolygonItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsPolygonItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsPolygonItem_SetFillRule(QtObjectPtr ptr, int rule);
void QGraphicsPolygonItem_SetPolygon(QtObjectPtr ptr, QtObjectPtr polygon);
int QGraphicsPolygonItem_Type(QtObjectPtr ptr);
void QGraphicsPolygonItem_DestroyQGraphicsPolygonItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif