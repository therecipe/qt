#ifdef __cplusplus
extern "C" {
#endif

void* QSizePolicy_NewQSizePolicy();
void* QSizePolicy_NewQSizePolicy2(int horizontal, int vertical, int ty);
int QSizePolicy_ControlType(void* ptr);
int QSizePolicy_ExpandingDirections(void* ptr);
int QSizePolicy_HasHeightForWidth(void* ptr);
int QSizePolicy_HasWidthForHeight(void* ptr);
int QSizePolicy_HorizontalPolicy(void* ptr);
int QSizePolicy_HorizontalStretch(void* ptr);
int QSizePolicy_RetainSizeWhenHidden(void* ptr);
void QSizePolicy_SetControlType(void* ptr, int ty);
void QSizePolicy_SetHeightForWidth(void* ptr, int dependent);
void QSizePolicy_SetHorizontalPolicy(void* ptr, int policy);
void QSizePolicy_SetHorizontalStretch(void* ptr, int stretchFactor);
void QSizePolicy_SetRetainSizeWhenHidden(void* ptr, int retainSize);
void QSizePolicy_SetVerticalPolicy(void* ptr, int policy);
void QSizePolicy_SetVerticalStretch(void* ptr, int stretchFactor);
void QSizePolicy_SetWidthForHeight(void* ptr, int dependent);
void QSizePolicy_Transpose(void* ptr);
int QSizePolicy_VerticalPolicy(void* ptr);
int QSizePolicy_VerticalStretch(void* ptr);

#ifdef __cplusplus
}
#endif