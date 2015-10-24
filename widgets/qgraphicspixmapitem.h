#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGraphicsPixmapItem_NewQGraphicsPixmapItem(QtObjectPtr parent);
QtObjectPtr QGraphicsPixmapItem_NewQGraphicsPixmapItem2(QtObjectPtr pixmap, QtObjectPtr parent);
int QGraphicsPixmapItem_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QGraphicsPixmapItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsPixmapItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsPixmapItem_SetOffset(QtObjectPtr ptr, QtObjectPtr offset);
void QGraphicsPixmapItem_SetPixmap(QtObjectPtr ptr, QtObjectPtr pixmap);
void QGraphicsPixmapItem_SetShapeMode(QtObjectPtr ptr, int mode);
void QGraphicsPixmapItem_SetTransformationMode(QtObjectPtr ptr, int mode);
int QGraphicsPixmapItem_ShapeMode(QtObjectPtr ptr);
int QGraphicsPixmapItem_TransformationMode(QtObjectPtr ptr);
int QGraphicsPixmapItem_Type(QtObjectPtr ptr);
void QGraphicsPixmapItem_DestroyQGraphicsPixmapItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif