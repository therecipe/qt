#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGraphicsSvgItem_NewQGraphicsSvgItem(QtObjectPtr parent);
QtObjectPtr QGraphicsSvgItem_NewQGraphicsSvgItem2(char* fileName, QtObjectPtr parent);
char* QGraphicsSvgItem_ElementId(QtObjectPtr ptr);
void QGraphicsSvgItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
QtObjectPtr QGraphicsSvgItem_Renderer(QtObjectPtr ptr);
void QGraphicsSvgItem_SetElementId(QtObjectPtr ptr, char* id);
void QGraphicsSvgItem_SetMaximumCacheSize(QtObjectPtr ptr, QtObjectPtr size);
void QGraphicsSvgItem_SetSharedRenderer(QtObjectPtr ptr, QtObjectPtr renderer);
int QGraphicsSvgItem_Type(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif