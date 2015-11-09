#ifdef __cplusplus
extern "C" {
#endif

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

#ifdef __cplusplus
}
#endif