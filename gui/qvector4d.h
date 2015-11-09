#ifdef __cplusplus
extern "C" {
#endif

void* QVector4D_NewQVector4D();
void* QVector4D_NewQVector4D4(void* point);
void* QVector4D_NewQVector4D5(void* point);
void* QVector4D_NewQVector4D6(void* vector);
void* QVector4D_NewQVector4D8(void* vector);
int QVector4D_IsNull(void* ptr);
void QVector4D_Normalize(void* ptr);

#ifdef __cplusplus
}
#endif