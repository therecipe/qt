#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSvgRenderer_FramesPerSecond(QtObjectPtr ptr);
void QSvgRenderer_SetFramesPerSecond(QtObjectPtr ptr, int num);
void QSvgRenderer_SetViewBox(QtObjectPtr ptr, QtObjectPtr viewbox);
void QSvgRenderer_SetViewBox2(QtObjectPtr ptr, QtObjectPtr viewbox);
QtObjectPtr QSvgRenderer_NewQSvgRenderer(QtObjectPtr parent);
QtObjectPtr QSvgRenderer_NewQSvgRenderer4(QtObjectPtr contents, QtObjectPtr parent);
QtObjectPtr QSvgRenderer_NewQSvgRenderer3(QtObjectPtr contents, QtObjectPtr parent);
QtObjectPtr QSvgRenderer_NewQSvgRenderer2(char* filename, QtObjectPtr parent);
int QSvgRenderer_Animated(QtObjectPtr ptr);
int QSvgRenderer_ElementExists(QtObjectPtr ptr, char* id);
int QSvgRenderer_IsValid(QtObjectPtr ptr);
int QSvgRenderer_Load3(QtObjectPtr ptr, QtObjectPtr contents);
int QSvgRenderer_Load2(QtObjectPtr ptr, QtObjectPtr contents);
int QSvgRenderer_Load(QtObjectPtr ptr, char* filename);
void QSvgRenderer_Render(QtObjectPtr ptr, QtObjectPtr painter);
void QSvgRenderer_Render2(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr bounds);
void QSvgRenderer_Render3(QtObjectPtr ptr, QtObjectPtr painter, char* elementId, QtObjectPtr bounds);
void QSvgRenderer_ConnectRepaintNeeded(QtObjectPtr ptr);
void QSvgRenderer_DisconnectRepaintNeeded(QtObjectPtr ptr);
void QSvgRenderer_DestroyQSvgRenderer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif