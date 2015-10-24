#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTransform_QTransform_QuadToSquare(QtObjectPtr quad, QtObjectPtr trans);
QtObjectPtr QTransform_NewQTransform();
int QTransform_IsAffine(QtObjectPtr ptr);
int QTransform_IsIdentity(QtObjectPtr ptr);
int QTransform_IsInvertible(QtObjectPtr ptr);
int QTransform_IsRotating(QtObjectPtr ptr);
int QTransform_IsScaling(QtObjectPtr ptr);
int QTransform_IsTranslating(QtObjectPtr ptr);
void QTransform_Map10(QtObjectPtr ptr, int x, int y, int tx, int ty);
int QTransform_QTransform_QuadToQuad(QtObjectPtr one, QtObjectPtr two, QtObjectPtr trans);
void QTransform_Reset(QtObjectPtr ptr);
int QTransform_QTransform_SquareToQuad(QtObjectPtr quad, QtObjectPtr trans);
int QTransform_Type(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif