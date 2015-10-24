#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QFrame_FrameShadow(QtObjectPtr ptr);
int QFrame_FrameShape(QtObjectPtr ptr);
int QFrame_FrameWidth(QtObjectPtr ptr);
int QFrame_LineWidth(QtObjectPtr ptr);
int QFrame_MidLineWidth(QtObjectPtr ptr);
void QFrame_SetFrameRect(QtObjectPtr ptr, QtObjectPtr v);
void QFrame_SetFrameShadow(QtObjectPtr ptr, int v);
void QFrame_SetFrameShape(QtObjectPtr ptr, int v);
void QFrame_SetLineWidth(QtObjectPtr ptr, int v);
void QFrame_SetMidLineWidth(QtObjectPtr ptr, int v);
QtObjectPtr QFrame_NewQFrame(QtObjectPtr parent, int f);
int QFrame_FrameStyle(QtObjectPtr ptr);
void QFrame_SetFrameStyle(QtObjectPtr ptr, int style);
void QFrame_DestroyQFrame(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif