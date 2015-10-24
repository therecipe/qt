#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QRect_Contains(QtObjectPtr ptr, QtObjectPtr point, int proper);
int QRect_Contains4(QtObjectPtr ptr, QtObjectPtr rectangle, int proper);
int QRect_Intersects(QtObjectPtr ptr, QtObjectPtr rectangle);
QtObjectPtr QRect_NewQRect();
QtObjectPtr QRect_NewQRect2(QtObjectPtr topLeft, QtObjectPtr bottomRight);
QtObjectPtr QRect_NewQRect3(QtObjectPtr topLeft, QtObjectPtr size);
QtObjectPtr QRect_NewQRect4(int x, int y, int width, int height);
void QRect_Adjust(QtObjectPtr ptr, int dx1, int dy1, int dx2, int dy2);
int QRect_Bottom(QtObjectPtr ptr);
int QRect_Contains3(QtObjectPtr ptr, int x, int y);
int QRect_Contains2(QtObjectPtr ptr, int x, int y, int proper);
void QRect_GetCoords(QtObjectPtr ptr, int x1, int y1, int x2, int y2);
void QRect_GetRect(QtObjectPtr ptr, int x, int y, int width, int height);
int QRect_Height(QtObjectPtr ptr);
int QRect_IsEmpty(QtObjectPtr ptr);
int QRect_IsNull(QtObjectPtr ptr);
int QRect_IsValid(QtObjectPtr ptr);
int QRect_Left(QtObjectPtr ptr);
void QRect_MoveBottom(QtObjectPtr ptr, int y);
void QRect_MoveBottomLeft(QtObjectPtr ptr, QtObjectPtr position);
void QRect_MoveBottomRight(QtObjectPtr ptr, QtObjectPtr position);
void QRect_MoveCenter(QtObjectPtr ptr, QtObjectPtr position);
void QRect_MoveLeft(QtObjectPtr ptr, int x);
void QRect_MoveRight(QtObjectPtr ptr, int x);
void QRect_MoveTo2(QtObjectPtr ptr, QtObjectPtr position);
void QRect_MoveTo(QtObjectPtr ptr, int x, int y);
void QRect_MoveTop(QtObjectPtr ptr, int y);
void QRect_MoveTopLeft(QtObjectPtr ptr, QtObjectPtr position);
void QRect_MoveTopRight(QtObjectPtr ptr, QtObjectPtr position);
int QRect_Right(QtObjectPtr ptr);
void QRect_SetBottom(QtObjectPtr ptr, int y);
void QRect_SetBottomLeft(QtObjectPtr ptr, QtObjectPtr position);
void QRect_SetBottomRight(QtObjectPtr ptr, QtObjectPtr position);
void QRect_SetCoords(QtObjectPtr ptr, int x1, int y1, int x2, int y2);
void QRect_SetHeight(QtObjectPtr ptr, int height);
void QRect_SetLeft(QtObjectPtr ptr, int x);
void QRect_SetRect(QtObjectPtr ptr, int x, int y, int width, int height);
void QRect_SetRight(QtObjectPtr ptr, int x);
void QRect_SetSize(QtObjectPtr ptr, QtObjectPtr size);
void QRect_SetTop(QtObjectPtr ptr, int y);
void QRect_SetTopLeft(QtObjectPtr ptr, QtObjectPtr position);
void QRect_SetTopRight(QtObjectPtr ptr, QtObjectPtr position);
void QRect_SetWidth(QtObjectPtr ptr, int width);
void QRect_SetX(QtObjectPtr ptr, int x);
void QRect_SetY(QtObjectPtr ptr, int y);
int QRect_Top(QtObjectPtr ptr);
void QRect_Translate2(QtObjectPtr ptr, QtObjectPtr offset);
void QRect_Translate(QtObjectPtr ptr, int dx, int dy);
int QRect_Width(QtObjectPtr ptr);
int QRect_X(QtObjectPtr ptr);
int QRect_Y(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif