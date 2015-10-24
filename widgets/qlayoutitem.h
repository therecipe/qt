#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QLayoutItem_Alignment(QtObjectPtr ptr);
int QLayoutItem_ControlTypes(QtObjectPtr ptr);
int QLayoutItem_ExpandingDirections(QtObjectPtr ptr);
int QLayoutItem_HasHeightForWidth(QtObjectPtr ptr);
int QLayoutItem_HeightForWidth(QtObjectPtr ptr, int w);
void QLayoutItem_Invalidate(QtObjectPtr ptr);
int QLayoutItem_IsEmpty(QtObjectPtr ptr);
QtObjectPtr QLayoutItem_Layout(QtObjectPtr ptr);
int QLayoutItem_MinimumHeightForWidth(QtObjectPtr ptr, int w);
void QLayoutItem_SetAlignment(QtObjectPtr ptr, int alignment);
void QLayoutItem_SetGeometry(QtObjectPtr ptr, QtObjectPtr r);
QtObjectPtr QLayoutItem_SpacerItem(QtObjectPtr ptr);
QtObjectPtr QLayoutItem_Widget(QtObjectPtr ptr);
void QLayoutItem_DestroyQLayoutItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif