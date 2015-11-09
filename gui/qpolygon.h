#ifdef __cplusplus
extern "C" {
#endif

void* QPolygon_NewQPolygon5(void* rectangle, int closed);
int QPolygon_ContainsPoint(void* ptr, void* point, int fillRule);
void QPolygon_PutPoints3(void* ptr, int index, int nPoints, void* fromPolygon, int fromIndex);
void* QPolygon_NewQPolygon();
void* QPolygon_NewQPolygon3(void* polygon);
void* QPolygon_NewQPolygon2(int size);
void QPolygon_Point(void* ptr, int index, int x, int y);
void QPolygon_SetPoint2(void* ptr, int index, void* point);
void QPolygon_SetPoint(void* ptr, int index, int x, int y);
void QPolygon_SetPoints(void* ptr, int nPoints, int points);
void QPolygon_Swap(void* ptr, void* other);
void QPolygon_Translate2(void* ptr, void* offset);
void QPolygon_Translate(void* ptr, int dx, int dy);
void QPolygon_DestroyQPolygon(void* ptr);

#ifdef __cplusplus
}
#endif