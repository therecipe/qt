#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSizePolicy_NewQSizePolicy();
QtObjectPtr QSizePolicy_NewQSizePolicy2(int horizontal, int vertical, int ty);
int QSizePolicy_ControlType(QtObjectPtr ptr);
int QSizePolicy_ExpandingDirections(QtObjectPtr ptr);
int QSizePolicy_HasHeightForWidth(QtObjectPtr ptr);
int QSizePolicy_HasWidthForHeight(QtObjectPtr ptr);
int QSizePolicy_HorizontalPolicy(QtObjectPtr ptr);
int QSizePolicy_HorizontalStretch(QtObjectPtr ptr);
int QSizePolicy_RetainSizeWhenHidden(QtObjectPtr ptr);
void QSizePolicy_SetControlType(QtObjectPtr ptr, int ty);
void QSizePolicy_SetHeightForWidth(QtObjectPtr ptr, int dependent);
void QSizePolicy_SetHorizontalPolicy(QtObjectPtr ptr, int policy);
void QSizePolicy_SetHorizontalStretch(QtObjectPtr ptr, int stretchFactor);
void QSizePolicy_SetRetainSizeWhenHidden(QtObjectPtr ptr, int retainSize);
void QSizePolicy_SetVerticalPolicy(QtObjectPtr ptr, int policy);
void QSizePolicy_SetVerticalStretch(QtObjectPtr ptr, int stretchFactor);
void QSizePolicy_SetWidthForHeight(QtObjectPtr ptr, int dependent);
void QSizePolicy_Transpose(QtObjectPtr ptr);
int QSizePolicy_VerticalPolicy(QtObjectPtr ptr);
int QSizePolicy_VerticalStretch(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif