#ifdef __cplusplus
extern "C" {
#endif

char* QGraphicsSvgItem_ElementId(void* ptr);
void QGraphicsSvgItem_Paint(void* ptr, void* painter, void* option, void* widget);
void* QGraphicsSvgItem_Renderer(void* ptr);
void QGraphicsSvgItem_SetElementId(void* ptr, char* id);
void QGraphicsSvgItem_SetMaximumCacheSize(void* ptr, void* size);
void QGraphicsSvgItem_SetSharedRenderer(void* ptr, void* renderer);
int QGraphicsSvgItem_Type(void* ptr);
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
char* QSvgGenerator_ObjectNameAbs(void* ptr);
void QSvgGenerator_SetObjectNameAbs(void* ptr, char* name);
int QSvgRenderer_FramesPerSecond(void* ptr);
void QSvgRenderer_SetFramesPerSecond(void* ptr, int num);
void QSvgRenderer_SetViewBox(void* ptr, void* viewbox);
void QSvgRenderer_SetViewBox2(void* ptr, void* viewbox);
void* QSvgRenderer_NewQSvgRenderer(void* parent);
void* QSvgRenderer_NewQSvgRenderer4(void* contents, void* parent);
void* QSvgRenderer_NewQSvgRenderer3(void* contents, void* parent);
void* QSvgRenderer_NewQSvgRenderer2(char* filename, void* parent);
int QSvgRenderer_Animated(void* ptr);
int QSvgRenderer_ElementExists(void* ptr, char* id);
int QSvgRenderer_IsValid(void* ptr);
int QSvgRenderer_Load3(void* ptr, void* contents);
int QSvgRenderer_Load2(void* ptr, void* contents);
int QSvgRenderer_Load(void* ptr, char* filename);
void QSvgRenderer_Render(void* ptr, void* painter);
void QSvgRenderer_Render2(void* ptr, void* painter, void* bounds);
void QSvgRenderer_Render3(void* ptr, void* painter, char* elementId, void* bounds);
void QSvgRenderer_ConnectRepaintNeeded(void* ptr);
void QSvgRenderer_DisconnectRepaintNeeded(void* ptr);
void QSvgRenderer_DestroyQSvgRenderer(void* ptr);
void* QSvgWidget_NewQSvgWidget(void* parent);
void* QSvgWidget_NewQSvgWidget2(char* file, void* parent);
void QSvgWidget_Load2(void* ptr, void* contents);
void QSvgWidget_Load(void* ptr, char* file);
void* QSvgWidget_Renderer(void* ptr);
void QSvgWidget_DestroyQSvgWidget(void* ptr);

#ifdef __cplusplus
}
#endif