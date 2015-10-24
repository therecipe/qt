#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSize_NewQSize();
QtObjectPtr QSize_NewQSize2(int width, int height);
int QSize_Height(QtObjectPtr ptr);
int QSize_IsEmpty(QtObjectPtr ptr);
int QSize_IsNull(QtObjectPtr ptr);
int QSize_IsValid(QtObjectPtr ptr);
int QSize_Rheight(QtObjectPtr ptr);
int QSize_Rwidth(QtObjectPtr ptr);
void QSize_Scale2(QtObjectPtr ptr, QtObjectPtr size, int mode);
void QSize_Scale(QtObjectPtr ptr, int width, int height, int mode);
void QSize_SetHeight(QtObjectPtr ptr, int height);
void QSize_SetWidth(QtObjectPtr ptr, int width);
void QSize_Transpose(QtObjectPtr ptr);
int QSize_Width(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif