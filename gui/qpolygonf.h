#ifdef __cplusplus
extern "C" {
#endif

void* QPolygonF_NewQPolygonF6(void* polygon);
void* QPolygonF_NewQPolygonF5(void* rectangle);
int QPolygonF_ContainsPoint(void* ptr, void* point, int fillRule);
void* QPolygonF_NewQPolygonF();
void* QPolygonF_NewQPolygonF3(void* polygon);
void* QPolygonF_NewQPolygonF2(int size);
int QPolygonF_IsClosed(void* ptr);
void QPolygonF_Swap(void* ptr, void* other);
void QPolygonF_Translate(void* ptr, void* offset);
void QPolygonF_Translate2(void* ptr, double dx, double dy);
void QPolygonF_DestroyQPolygonF(void* ptr);

#ifdef __cplusplus
}
#endif