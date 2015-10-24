#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGraphicsProxyWidget_NewQGraphicsProxyWidget(QtObjectPtr parent, int wFlags);
QtObjectPtr QGraphicsProxyWidget_CreateProxyForChildWidget(QtObjectPtr ptr, QtObjectPtr child);
void QGraphicsProxyWidget_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsProxyWidget_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect);
void QGraphicsProxyWidget_SetWidget(QtObjectPtr ptr, QtObjectPtr widget);
int QGraphicsProxyWidget_Type(QtObjectPtr ptr);
QtObjectPtr QGraphicsProxyWidget_Widget(QtObjectPtr ptr);
void QGraphicsProxyWidget_DestroyQGraphicsProxyWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif