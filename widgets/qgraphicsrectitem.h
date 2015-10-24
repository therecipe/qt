#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QGraphicsRectItem_SetRect(QtObjectPtr ptr, QtObjectPtr rectangle);
int QGraphicsRectItem_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QGraphicsRectItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsRectItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
int QGraphicsRectItem_Type(QtObjectPtr ptr);
void QGraphicsRectItem_DestroyQGraphicsRectItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif