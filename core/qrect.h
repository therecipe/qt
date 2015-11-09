#ifdef __cplusplus
extern "C" {
#endif

int QRect_Contains(void* ptr, void* point, int proper);
int QRect_Contains4(void* ptr, void* rectangle, int proper);
int QRect_Intersects(void* ptr, void* rectangle);
void* QRect_NewQRect();
void* QRect_NewQRect2(void* topLeft, void* bottomRight);
void* QRect_NewQRect3(void* topLeft, void* size);
void* QRect_NewQRect4(int x, int y, int width, int height);
void QRect_Adjust(void* ptr, int dx1, int dy1, int dx2, int dy2);
int QRect_Bottom(void* ptr);
int QRect_Contains3(void* ptr, int x, int y);
int QRect_Contains2(void* ptr, int x, int y, int proper);
void QRect_GetCoords(void* ptr, int x1, int y1, int x2, int y2);
void QRect_GetRect(void* ptr, int x, int y, int width, int height);
int QRect_Height(void* ptr);
int QRect_IsEmpty(void* ptr);
int QRect_IsNull(void* ptr);
int QRect_IsValid(void* ptr);
int QRect_Left(void* ptr);
void QRect_MoveBottom(void* ptr, int y);
void QRect_MoveBottomLeft(void* ptr, void* position);
void QRect_MoveBottomRight(void* ptr, void* position);
void QRect_MoveCenter(void* ptr, void* position);
void QRect_MoveLeft(void* ptr, int x);
void QRect_MoveRight(void* ptr, int x);
void QRect_MoveTo2(void* ptr, void* position);
void QRect_MoveTo(void* ptr, int x, int y);
void QRect_MoveTop(void* ptr, int y);
void QRect_MoveTopLeft(void* ptr, void* position);
void QRect_MoveTopRight(void* ptr, void* position);
int QRect_Right(void* ptr);
void QRect_SetBottom(void* ptr, int y);
void QRect_SetBottomLeft(void* ptr, void* position);
void QRect_SetBottomRight(void* ptr, void* position);
void QRect_SetCoords(void* ptr, int x1, int y1, int x2, int y2);
void QRect_SetHeight(void* ptr, int height);
void QRect_SetLeft(void* ptr, int x);
void QRect_SetRect(void* ptr, int x, int y, int width, int height);
void QRect_SetRight(void* ptr, int x);
void QRect_SetSize(void* ptr, void* size);
void QRect_SetTop(void* ptr, int y);
void QRect_SetTopLeft(void* ptr, void* position);
void QRect_SetTopRight(void* ptr, void* position);
void QRect_SetWidth(void* ptr, int width);
void QRect_SetX(void* ptr, int x);
void QRect_SetY(void* ptr, int y);
int QRect_Top(void* ptr);
void QRect_Translate2(void* ptr, void* offset);
void QRect_Translate(void* ptr, int dx, int dy);
int QRect_Width(void* ptr);
int QRect_X(void* ptr);
int QRect_Y(void* ptr);

#ifdef __cplusplus
}
#endif