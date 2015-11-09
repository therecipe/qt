#ifdef __cplusplus
extern "C" {
#endif

void* QVector3D_NewQVector3D();
void* QVector3D_NewQVector3D4(void* point);
void* QVector3D_NewQVector3D5(void* point);
void* QVector3D_NewQVector3D6(void* vector);
void* QVector3D_NewQVector3D8(void* vector);
int QVector3D_IsNull(void* ptr);
void QVector3D_Normalize(void* ptr);

#ifdef __cplusplus
}
#endif