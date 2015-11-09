#ifdef __cplusplus
extern "C" {
#endif

void* QMatrix4x4_NewQMatrix4x4();
void* QMatrix4x4_NewQMatrix4x47(void* transform);
int QMatrix4x4_IsAffine(void* ptr);
int QMatrix4x4_IsIdentity(void* ptr);
void QMatrix4x4_LookAt(void* ptr, void* eye, void* center, void* up);
void QMatrix4x4_Optimize(void* ptr);
void QMatrix4x4_Ortho2(void* ptr, void* rect);
void QMatrix4x4_Ortho3(void* ptr, void* rect);
void QMatrix4x4_Rotate2(void* ptr, void* quaternion);
void QMatrix4x4_Scale(void* ptr, void* vector);
void QMatrix4x4_SetColumn(void* ptr, int index, void* value);
void QMatrix4x4_SetRow(void* ptr, int index, void* value);
void QMatrix4x4_SetToIdentity(void* ptr);
void QMatrix4x4_Translate(void* ptr, void* vector);
void QMatrix4x4_Viewport2(void* ptr, void* rect);

#ifdef __cplusplus
}
#endif