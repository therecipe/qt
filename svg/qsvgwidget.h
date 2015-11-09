#ifdef __cplusplus
extern "C" {
#endif

void* QSvgWidget_NewQSvgWidget(void* parent);
void* QSvgWidget_NewQSvgWidget2(char* file, void* parent);
void QSvgWidget_Load2(void* ptr, void* contents);
void QSvgWidget_Load(void* ptr, char* file);
void* QSvgWidget_Renderer(void* ptr);
void QSvgWidget_DestroyQSvgWidget(void* ptr);

#ifdef __cplusplus
}
#endif