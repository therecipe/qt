#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGraphicsLineItem_NewQGraphicsLineItem(QtObjectPtr parent);
QtObjectPtr QGraphicsLineItem_NewQGraphicsLineItem2(QtObjectPtr line, QtObjectPtr parent);
int QGraphicsLineItem_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QGraphicsLineItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsLineItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsLineItem_SetLine(QtObjectPtr ptr, QtObjectPtr line);
void QGraphicsLineItem_SetPen(QtObjectPtr ptr, QtObjectPtr pen);
int QGraphicsLineItem_Type(QtObjectPtr ptr);
void QGraphicsLineItem_DestroyQGraphicsLineItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif