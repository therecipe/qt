#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QWidgetItem_NewQWidgetItem(QtObjectPtr widget);
int QWidgetItem_ControlTypes(QtObjectPtr ptr);
int QWidgetItem_ExpandingDirections(QtObjectPtr ptr);
int QWidgetItem_HasHeightForWidth(QtObjectPtr ptr);
int QWidgetItem_HeightForWidth(QtObjectPtr ptr, int w);
int QWidgetItem_IsEmpty(QtObjectPtr ptr);
void QWidgetItem_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect);
QtObjectPtr QWidgetItem_Widget(QtObjectPtr ptr);
void QWidgetItem_DestroyQWidgetItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif