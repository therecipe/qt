#ifdef __cplusplus
extern "C" {
#endif

void* QVector2D_NewQVector2D();
void* QVector2D_NewQVector2D4(void* point);
void* QVector2D_NewQVector2D5(void* point);
void* QVector2D_NewQVector2D6(void* vector);
void* QVector2D_NewQVector2D7(void* vector);
int QVector2D_IsNull(void* ptr);
void QVector2D_Normalize(void* ptr);

#ifdef __cplusplus
}
#endif