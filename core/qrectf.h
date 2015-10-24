#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QRectF_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QRectF_Contains3(QtObjectPtr ptr, QtObjectPtr rectangle);
int QRectF_Intersects(QtObjectPtr ptr, QtObjectPtr rectangle);
QtObjectPtr QRectF_NewQRectF();
QtObjectPtr QRectF_NewQRectF3(QtObjectPtr topLeft, QtObjectPtr bottomRight);
QtObjectPtr QRectF_NewQRectF2(QtObjectPtr topLeft, QtObjectPtr size);
QtObjectPtr QRectF_NewQRectF5(QtObjectPtr rectangle);
int QRectF_IsEmpty(QtObjectPtr ptr);
int QRectF_IsNull(QtObjectPtr ptr);
int QRectF_IsValid(QtObjectPtr ptr);
void QRectF_MoveBottomLeft(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_MoveBottomRight(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_MoveCenter(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_MoveTo2(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_MoveTopLeft(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_MoveTopRight(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_SetBottomLeft(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_SetBottomRight(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_SetSize(QtObjectPtr ptr, QtObjectPtr size);
void QRectF_SetTopLeft(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_SetTopRight(QtObjectPtr ptr, QtObjectPtr position);
void QRectF_Translate2(QtObjectPtr ptr, QtObjectPtr offset);

#ifdef __cplusplus
}
#endif