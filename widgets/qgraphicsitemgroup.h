#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGraphicsItemGroup_NewQGraphicsItemGroup(QtObjectPtr parent);
void QGraphicsItemGroup_AddToGroup(QtObjectPtr ptr, QtObjectPtr item);
int QGraphicsItemGroup_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsItemGroup_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsItemGroup_RemoveFromGroup(QtObjectPtr ptr, QtObjectPtr item);
int QGraphicsItemGroup_Type(QtObjectPtr ptr);
void QGraphicsItemGroup_DestroyQGraphicsItemGroup(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif