#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGraphicsPathItem_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QGraphicsPathItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsPathItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsPathItem_SetPath(QtObjectPtr ptr, QtObjectPtr path);
int QGraphicsPathItem_Type(QtObjectPtr ptr);
void QGraphicsPathItem_DestroyQGraphicsPathItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif