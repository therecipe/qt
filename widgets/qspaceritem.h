#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSpacerItem_NewQSpacerItem(int w, int h, int hPolicy, int vPolicy);
void QSpacerItem_ChangeSize(QtObjectPtr ptr, int w, int h, int hPolicy, int vPolicy);
int QSpacerItem_ExpandingDirections(QtObjectPtr ptr);
int QSpacerItem_IsEmpty(QtObjectPtr ptr);
void QSpacerItem_SetGeometry(QtObjectPtr ptr, QtObjectPtr r);
QtObjectPtr QSpacerItem_SpacerItem(QtObjectPtr ptr);
void QSpacerItem_DestroyQSpacerItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif