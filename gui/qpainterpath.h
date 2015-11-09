#ifdef __cplusplus
extern "C" {
#endif

void* QPainterPath_NewQPainterPath3(void* path);
void QPainterPath_AddEllipse(void* ptr, void* boundingRectangle);
void QPainterPath_AddPath(void* ptr, void* path);
void QPainterPath_AddRect(void* ptr, void* rectangle);
void QPainterPath_AddText(void* ptr, void* point, void* font, char* text);
void QPainterPath_ArcMoveTo(void* ptr, void* rectangle, double angle);
void QPainterPath_ArcTo(void* ptr, void* rectangle, double startAngle, double sweepLength);
void QPainterPath_ConnectPath(void* ptr, void* path);
int QPainterPath_Contains(void* ptr, void* point);
int QPainterPath_Contains2(void* ptr, void* rectangle);
void QPainterPath_CubicTo(void* ptr, void* c1, void* c2, void* endPoint);
int QPainterPath_ElementCount(void* ptr);
int QPainterPath_Intersects(void* ptr, void* rectangle);
int QPainterPath_IsEmpty(void* ptr);
void QPainterPath_LineTo(void* ptr, void* endPoint);
void QPainterPath_MoveTo(void* ptr, void* point);
void QPainterPath_QuadTo(void* ptr, void* c, void* endPoint);
void QPainterPath_SetElementPositionAt(void* ptr, int index, double x, double y);
void QPainterPath_SetFillRule(void* ptr, int fillRule);
void* QPainterPath_NewQPainterPath();
void* QPainterPath_NewQPainterPath2(void* startPoint);
void QPainterPath_AddEllipse3(void* ptr, void* center, double rx, double ry);
void QPainterPath_AddEllipse2(void* ptr, double x, double y, double width, double height);
void QPainterPath_AddPolygon(void* ptr, void* polygon);
void QPainterPath_AddRect2(void* ptr, double x, double y, double width, double height);
void QPainterPath_AddRegion(void* ptr, void* region);
void QPainterPath_AddRoundedRect(void* ptr, void* rect, double xRadius, double yRadius, int mode);
void QPainterPath_AddRoundedRect2(void* ptr, double x, double y, double w, double h, double xRadius, double yRadius, int mode);
void QPainterPath_AddText2(void* ptr, double x, double y, void* font, char* text);
double QPainterPath_AngleAtPercent(void* ptr, double t);
void QPainterPath_ArcMoveTo2(void* ptr, double x, double y, double width, double height, double angle);
void QPainterPath_ArcTo2(void* ptr, double x, double y, double width, double height, double startAngle, double sweepLength);
void QPainterPath_CloseSubpath(void* ptr);
int QPainterPath_Contains3(void* ptr, void* p);
void QPainterPath_CubicTo2(void* ptr, double c1X, double c1Y, double c2X, double c2Y, double endPointX, double endPointY);
int QPainterPath_FillRule(void* ptr);
int QPainterPath_Intersects2(void* ptr, void* p);
double QPainterPath_Length(void* ptr);
void QPainterPath_LineTo2(void* ptr, double x, double y);
void QPainterPath_MoveTo2(void* ptr, double x, double y);
double QPainterPath_PercentAtLength(void* ptr, double len);
void QPainterPath_QuadTo2(void* ptr, double cx, double cy, double endPointX, double endPointY);
double QPainterPath_SlopeAtPercent(void* ptr, double t);
void QPainterPath_Swap(void* ptr, void* other);
void QPainterPath_Translate2(void* ptr, void* offset);
void QPainterPath_Translate(void* ptr, double dx, double dy);
void QPainterPath_DestroyQPainterPath(void* ptr);

#ifdef __cplusplus
}
#endif