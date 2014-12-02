#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
void QRect_Adjust_Int_Int_Int_Int(QtObjectPtr ptr, int dx1, int dy1, int dx2, int dy2);
int QRect_Bottom(QtObjectPtr ptr);
int QRect_Contains_Int_Int_Bool(QtObjectPtr ptr, int x, int y, int proper);
int QRect_Contains_Int_Int(QtObjectPtr ptr, int x, int y);
int QRect_Height(QtObjectPtr ptr);
int QRect_IsEmpty(QtObjectPtr ptr);
int QRect_IsNull(QtObjectPtr ptr);
int QRect_IsValid(QtObjectPtr ptr);
int QRect_Left(QtObjectPtr ptr);
void QRect_MoveBottom_Int(QtObjectPtr ptr, int y);
void QRect_MoveLeft_Int(QtObjectPtr ptr, int x);
void QRect_MoveRight_Int(QtObjectPtr ptr, int x);
void QRect_MoveTo_Int_Int(QtObjectPtr ptr, int x, int y);
void QRect_MoveTop_Int(QtObjectPtr ptr, int y);
int QRect_Right(QtObjectPtr ptr);
void QRect_SetBottom_Int(QtObjectPtr ptr, int y);
void QRect_SetCoords_Int_Int_Int_Int(QtObjectPtr ptr, int x1, int y1, int x2, int y2);
void QRect_SetHeight_Int(QtObjectPtr ptr, int height);
void QRect_SetLeft_Int(QtObjectPtr ptr, int x);
void QRect_SetRect_Int_Int_Int_Int(QtObjectPtr ptr, int x, int y, int width, int height);
void QRect_SetRight_Int(QtObjectPtr ptr, int x);
void QRect_SetTop_Int(QtObjectPtr ptr, int y);
void QRect_SetWidth_Int(QtObjectPtr ptr, int width);
void QRect_SetX_Int(QtObjectPtr ptr, int x);
void QRect_SetY_Int(QtObjectPtr ptr, int y);
int QRect_Top(QtObjectPtr ptr);
void QRect_Translate_Int_Int(QtObjectPtr ptr, int dx, int dy);
int QRect_Width(QtObjectPtr ptr);
int QRect_X(QtObjectPtr ptr);
int QRect_Y(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
