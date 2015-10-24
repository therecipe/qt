#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMargins_NewQMargins();
QtObjectPtr QMargins_NewQMargins2(int left, int top, int right, int bottom);
int QMargins_Bottom(QtObjectPtr ptr);
int QMargins_IsNull(QtObjectPtr ptr);
int QMargins_Left(QtObjectPtr ptr);
int QMargins_Right(QtObjectPtr ptr);
void QMargins_SetBottom(QtObjectPtr ptr, int bottom);
void QMargins_SetLeft(QtObjectPtr ptr, int left);
void QMargins_SetRight(QtObjectPtr ptr, int right);
void QMargins_SetTop(QtObjectPtr ptr, int Top);
int QMargins_Top(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif