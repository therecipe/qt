#ifdef __cplusplus
extern "C" {
#endif

void* QPoint_NewQPoint();
void* QPoint_NewQPoint2(int xpos, int ypos);
int QPoint_QPoint_DotProduct(void* p1, void* p2);
int QPoint_IsNull(void* ptr);
int QPoint_ManhattanLength(void* ptr);
int QPoint_Rx(void* ptr);
int QPoint_Ry(void* ptr);
void QPoint_SetX(void* ptr, int x);
void QPoint_SetY(void* ptr, int y);
int QPoint_X(void* ptr);
int QPoint_Y(void* ptr);

#ifdef __cplusplus
}
#endif