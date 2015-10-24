#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPainterPath_NewQPainterPath3(QtObjectPtr path);
void QPainterPath_AddEllipse(QtObjectPtr ptr, QtObjectPtr boundingRectangle);
void QPainterPath_AddPath(QtObjectPtr ptr, QtObjectPtr path);
void QPainterPath_AddRect(QtObjectPtr ptr, QtObjectPtr rectangle);
void QPainterPath_AddText(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr font, char* text);
void QPainterPath_ConnectPath(QtObjectPtr ptr, QtObjectPtr path);
int QPainterPath_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QPainterPath_Contains2(QtObjectPtr ptr, QtObjectPtr rectangle);
void QPainterPath_CubicTo(QtObjectPtr ptr, QtObjectPtr c1, QtObjectPtr c2, QtObjectPtr endPoint);
int QPainterPath_ElementCount(QtObjectPtr ptr);
int QPainterPath_Intersects(QtObjectPtr ptr, QtObjectPtr rectangle);
int QPainterPath_IsEmpty(QtObjectPtr ptr);
void QPainterPath_LineTo(QtObjectPtr ptr, QtObjectPtr endPoint);
void QPainterPath_MoveTo(QtObjectPtr ptr, QtObjectPtr point);
void QPainterPath_QuadTo(QtObjectPtr ptr, QtObjectPtr c, QtObjectPtr endPoint);
void QPainterPath_SetFillRule(QtObjectPtr ptr, int fillRule);
QtObjectPtr QPainterPath_NewQPainterPath();
QtObjectPtr QPainterPath_NewQPainterPath2(QtObjectPtr startPoint);
void QPainterPath_AddPolygon(QtObjectPtr ptr, QtObjectPtr polygon);
void QPainterPath_AddRegion(QtObjectPtr ptr, QtObjectPtr region);
void QPainterPath_CloseSubpath(QtObjectPtr ptr);
int QPainterPath_Contains3(QtObjectPtr ptr, QtObjectPtr p);
int QPainterPath_FillRule(QtObjectPtr ptr);
int QPainterPath_Intersects2(QtObjectPtr ptr, QtObjectPtr p);
void QPainterPath_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QPainterPath_Translate2(QtObjectPtr ptr, QtObjectPtr offset);
void QPainterPath_DestroyQPainterPath(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif