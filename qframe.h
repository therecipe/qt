#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QFrame_New_QWidget_WindowType(QtObjectPtr parent, int f);
void QFrame_Destroy(QtObjectPtr ptr);
int QFrame_FrameStyle(QtObjectPtr ptr);
int QFrame_FrameWidth(QtObjectPtr ptr);
int QFrame_LineWidth(QtObjectPtr ptr);
int QFrame_MidLineWidth(QtObjectPtr ptr);
void QFrame_SetFrameStyle_Int(QtObjectPtr ptr, int style);
void QFrame_SetLineWidth_Int(QtObjectPtr ptr, int width);
void QFrame_SetMidLineWidth_Int(QtObjectPtr ptr, int width);

#ifdef __cplusplus
}
#endif
