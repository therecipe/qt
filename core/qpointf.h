#ifdef __cplusplus
extern "C" {
#endif

void* QPointF_NewQPointF();
void* QPointF_NewQPointF2(void* point);
void* QPointF_NewQPointF3(double xpos, double ypos);
double QPointF_QPointF_DotProduct(void* p1, void* p2);
int QPointF_IsNull(void* ptr);
double QPointF_ManhattanLength(void* ptr);
double QPointF_Rx(void* ptr);
double QPointF_Ry(void* ptr);
void QPointF_SetX(void* ptr, double x);
void QPointF_SetY(void* ptr, double y);
double QPointF_X(void* ptr);
double QPointF_Y(void* ptr);

#ifdef __cplusplus
}
#endif