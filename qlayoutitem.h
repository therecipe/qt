#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
void QLayoutItem_Destroy(QtObjectPtr ptr);
int QLayoutItem_Alignment(QtObjectPtr ptr);
int QLayoutItem_ExpandingDirections(QtObjectPtr ptr);
int QLayoutItem_HasHeightForWidth(QtObjectPtr ptr);
int QLayoutItem_HeightForWidth_Int(QtObjectPtr ptr, int w);
void QLayoutItem_Invalidate(QtObjectPtr ptr);
int QLayoutItem_IsEmpty(QtObjectPtr ptr);
QtObjectPtr QLayoutItem_Layout(QtObjectPtr ptr);
int QLayoutItem_MinimumHeightForWidth_Int(QtObjectPtr ptr, int w);
void QLayoutItem_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment);
QtObjectPtr QLayoutItem_Widget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
