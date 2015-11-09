#ifdef __cplusplus
extern "C" {
#endif

void QIconEngine_AddFile(void* ptr, char* fileName, void* size, int mode, int state);
void QIconEngine_AddPixmap(void* ptr, void* pixmap, int mode, int state);
void* QIconEngine_Clone(void* ptr);
char* QIconEngine_IconName(void* ptr);
char* QIconEngine_Key(void* ptr);
void QIconEngine_Paint(void* ptr, void* painter, void* rect, int mode, int state);
int QIconEngine_Read(void* ptr, void* in);
int QIconEngine_Write(void* ptr, void* out);
void QIconEngine_DestroyQIconEngine(void* ptr);

#ifdef __cplusplus
}
#endif