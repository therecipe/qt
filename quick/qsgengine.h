#ifdef __cplusplus
extern "C" {
#endif

void* QSGEngine_NewQSGEngine(void* parent);
void* QSGEngine_CreateRenderer(void* ptr);
void* QSGEngine_CreateTextureFromImage(void* ptr, void* image, int options);
void QSGEngine_Initialize(void* ptr, void* context);
void QSGEngine_Invalidate(void* ptr);
void QSGEngine_DestroyQSGEngine(void* ptr);

#ifdef __cplusplus
}
#endif