#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QPaintEngine_DrawEllipse(QtObjectPtr ptr, QtObjectPtr rect);
void QPaintEngine_DrawImage(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr image, QtObjectPtr sr, int flags);
void QPaintEngine_DrawPolygon(QtObjectPtr ptr, QtObjectPtr points, int pointCount, int mode);
int QPaintEngine_Begin(QtObjectPtr ptr, QtObjectPtr pdev);
void QPaintEngine_DrawEllipse2(QtObjectPtr ptr, QtObjectPtr rect);
void QPaintEngine_DrawLines2(QtObjectPtr ptr, QtObjectPtr lines, int lineCount);
void QPaintEngine_DrawLines(QtObjectPtr ptr, QtObjectPtr lines, int lineCount);
void QPaintEngine_DrawPath(QtObjectPtr ptr, QtObjectPtr path);
void QPaintEngine_DrawPixmap(QtObjectPtr ptr, QtObjectPtr r, QtObjectPtr pm, QtObjectPtr sr);
void QPaintEngine_DrawPoints2(QtObjectPtr ptr, QtObjectPtr points, int pointCount);
void QPaintEngine_DrawPoints(QtObjectPtr ptr, QtObjectPtr points, int pointCount);
void QPaintEngine_DrawPolygon2(QtObjectPtr ptr, QtObjectPtr points, int pointCount, int mode);
void QPaintEngine_DrawRects2(QtObjectPtr ptr, QtObjectPtr rects, int rectCount);
void QPaintEngine_DrawRects(QtObjectPtr ptr, QtObjectPtr rects, int rectCount);
void QPaintEngine_DrawTextItem(QtObjectPtr ptr, QtObjectPtr p, QtObjectPtr textItem);
void QPaintEngine_DrawTiledPixmap(QtObjectPtr ptr, QtObjectPtr rect, QtObjectPtr pixmap, QtObjectPtr p);
int QPaintEngine_End(QtObjectPtr ptr);
int QPaintEngine_HasFeature(QtObjectPtr ptr, int feature);
int QPaintEngine_IsActive(QtObjectPtr ptr);
QtObjectPtr QPaintEngine_PaintDevice(QtObjectPtr ptr);
QtObjectPtr QPaintEngine_Painter(QtObjectPtr ptr);
void QPaintEngine_SetActive(QtObjectPtr ptr, int state);
int QPaintEngine_Type(QtObjectPtr ptr);
void QPaintEngine_UpdateState(QtObjectPtr ptr, QtObjectPtr state);
void QPaintEngine_DestroyQPaintEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif