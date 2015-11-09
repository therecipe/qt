#ifdef __cplusplus
extern "C" {
#endif

void* QTransform_NewQTransform3(double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33);
void* QTransform_NewQTransform4(double m11, double m12, double m21, double m22, double dx, double dy);
void* QTransform_Map8(void* ptr, void* region);
int QTransform_QTransform_QuadToSquare(void* quad, void* trans);
void* QTransform_NewQTransform();
double QTransform_Determinant(void* ptr);
double QTransform_Dx(void* ptr);
double QTransform_Dy(void* ptr);
int QTransform_IsAffine(void* ptr);
int QTransform_IsIdentity(void* ptr);
int QTransform_IsInvertible(void* ptr);
int QTransform_IsRotating(void* ptr);
int QTransform_IsScaling(void* ptr);
int QTransform_IsTranslating(void* ptr);
double QTransform_M11(void* ptr);
double QTransform_M12(void* ptr);
double QTransform_M13(void* ptr);
double QTransform_M21(void* ptr);
double QTransform_M22(void* ptr);
double QTransform_M23(void* ptr);
double QTransform_M31(void* ptr);
double QTransform_M32(void* ptr);
double QTransform_M33(void* ptr);
void QTransform_Map10(void* ptr, int x, int y, int tx, int ty);
int QTransform_QTransform_QuadToQuad(void* one, void* two, void* trans);
void QTransform_Reset(void* ptr);
void QTransform_SetMatrix(void* ptr, double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33);
int QTransform_QTransform_SquareToQuad(void* quad, void* trans);
int QTransform_Type(void* ptr);

#ifdef __cplusplus
}
#endif