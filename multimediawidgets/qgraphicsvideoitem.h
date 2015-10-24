#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGraphicsVideoItem_NewQGraphicsVideoItem(QtObjectPtr parent);
int QGraphicsVideoItem_AspectRatioMode(QtObjectPtr ptr);
QtObjectPtr QGraphicsVideoItem_MediaObject(QtObjectPtr ptr);
void QGraphicsVideoItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsVideoItem_SetAspectRatioMode(QtObjectPtr ptr, int mode);
void QGraphicsVideoItem_SetOffset(QtObjectPtr ptr, QtObjectPtr offset);
void QGraphicsVideoItem_SetSize(QtObjectPtr ptr, QtObjectPtr size);
void QGraphicsVideoItem_DestroyQGraphicsVideoItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif