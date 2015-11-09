#ifdef __cplusplus
extern "C" {
#endif

void QPaintEngine_DrawEllipse(void* ptr, void* rect);
void QPaintEngine_DrawImage(void* ptr, void* rectangle, void* image, void* sr, int flags);
void QPaintEngine_DrawPolygon(void* ptr, void* points, int pointCount, int mode);
int QPaintEngine_Begin(void* ptr, void* pdev);
void QPaintEngine_DrawEllipse2(void* ptr, void* rect);
void QPaintEngine_DrawLines2(void* ptr, void* lines, int lineCount);
void QPaintEngine_DrawLines(void* ptr, void* lines, int lineCount);
void QPaintEngine_DrawPath(void* ptr, void* path);
void QPaintEngine_DrawPixmap(void* ptr, void* r, void* pm, void* sr);
void QPaintEngine_DrawPoints2(void* ptr, void* points, int pointCount);
void QPaintEngine_DrawPoints(void* ptr, void* points, int pointCount);
void QPaintEngine_DrawPolygon2(void* ptr, void* points, int pointCount, int mode);
void QPaintEngine_DrawRects2(void* ptr, void* rects, int rectCount);
void QPaintEngine_DrawRects(void* ptr, void* rects, int rectCount);
void QPaintEngine_DrawTextItem(void* ptr, void* p, void* textItem);
void QPaintEngine_DrawTiledPixmap(void* ptr, void* rect, void* pixmap, void* p);
int QPaintEngine_End(void* ptr);
int QPaintEngine_HasFeature(void* ptr, int feature);
int QPaintEngine_IsActive(void* ptr);
void* QPaintEngine_PaintDevice(void* ptr);
void* QPaintEngine_Painter(void* ptr);
void QPaintEngine_SetActive(void* ptr, int state);
int QPaintEngine_Type(void* ptr);
void QPaintEngine_UpdateState(void* ptr, void* state);
void QPaintEngine_DestroyQPaintEngine(void* ptr);

#ifdef __cplusplus
}
#endif