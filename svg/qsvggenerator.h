#ifdef __cplusplus
extern "C" {
#endif

char* QSvgGenerator_Description(void* ptr);
char* QSvgGenerator_FileName(void* ptr);
void* QSvgGenerator_OutputDevice(void* ptr);
int QSvgGenerator_Resolution(void* ptr);
void QSvgGenerator_SetDescription(void* ptr, char* description);
void QSvgGenerator_SetFileName(void* ptr, char* fileName);
void QSvgGenerator_SetOutputDevice(void* ptr, void* outputDevice);
void QSvgGenerator_SetResolution(void* ptr, int dpi);
void QSvgGenerator_SetSize(void* ptr, void* size);
void QSvgGenerator_SetTitle(void* ptr, char* title);
void QSvgGenerator_SetViewBox(void* ptr, void* viewBox);
void QSvgGenerator_SetViewBox2(void* ptr, void* viewBox);
char* QSvgGenerator_Title(void* ptr);
void* QSvgGenerator_NewQSvgGenerator();
void QSvgGenerator_DestroyQSvgGenerator(void* ptr);

#ifdef __cplusplus
}
#endif