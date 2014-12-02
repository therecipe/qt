#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
int QSize_Height(QtObjectPtr ptr);
int QSize_IsEmpty(QtObjectPtr ptr);
int QSize_IsNull(QtObjectPtr ptr);
int QSize_IsValid(QtObjectPtr ptr);
int QSize_Rheight(QtObjectPtr ptr);
int QSize_Rwidth(QtObjectPtr ptr);
void QSize_Scale_Int_Int_AspectRatioMode(QtObjectPtr ptr, int width, int height, int mode);
void QSize_SetHeight_Int(QtObjectPtr ptr, int height);
void QSize_SetWidth_Int(QtObjectPtr ptr, int width);
void QSize_Transpose(QtObjectPtr ptr);
int QSize_Width(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
