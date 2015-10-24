#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGraphicsEllipseItem_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QGraphicsEllipseItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsEllipseItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsEllipseItem_SetRect(QtObjectPtr ptr, QtObjectPtr rect);
void QGraphicsEllipseItem_SetSpanAngle(QtObjectPtr ptr, int angle);
void QGraphicsEllipseItem_SetStartAngle(QtObjectPtr ptr, int angle);
int QGraphicsEllipseItem_SpanAngle(QtObjectPtr ptr);
int QGraphicsEllipseItem_StartAngle(QtObjectPtr ptr);
int QGraphicsEllipseItem_Type(QtObjectPtr ptr);
void QGraphicsEllipseItem_DestroyQGraphicsEllipseItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif