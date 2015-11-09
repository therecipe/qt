#ifdef __cplusplus
extern "C" {
#endif

void* QQuaternion_NewQQuaternion();
void* QQuaternion_NewQQuaternion5(void* vector);
void QQuaternion_GetAxes(void* ptr, void* xAxis, void* yAxis, void* zAxis);
int QQuaternion_IsIdentity(void* ptr);
int QQuaternion_IsNull(void* ptr);
void QQuaternion_Normalize(void* ptr);
void QQuaternion_SetVector(void* ptr, void* vector);

#ifdef __cplusplus
}
#endif