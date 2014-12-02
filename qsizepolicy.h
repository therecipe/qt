#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QSizePolicy_New();
int QSizePolicy_ExpandingDirections(QtObjectPtr ptr);
int QSizePolicy_HasHeightForWidth(QtObjectPtr ptr);
int QSizePolicy_HasWidthForHeight(QtObjectPtr ptr);
int QSizePolicy_HorizontalStretch(QtObjectPtr ptr);
int QSizePolicy_RetainSizeWhenHidden(QtObjectPtr ptr);
void QSizePolicy_SetHeightForWidth_Bool(QtObjectPtr ptr, int dependent);
void QSizePolicy_SetHorizontalStretch_Int(QtObjectPtr ptr, int stretchFactor);
void QSizePolicy_SetRetainSizeWhenHidden_Bool(QtObjectPtr ptr, int retainSize);
void QSizePolicy_SetVerticalStretch_Int(QtObjectPtr ptr, int stretchFactor);
void QSizePolicy_SetWidthForHeight_Bool(QtObjectPtr ptr, int dependent);
void QSizePolicy_Transpose(QtObjectPtr ptr);
int QSizePolicy_VerticalStretch(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
