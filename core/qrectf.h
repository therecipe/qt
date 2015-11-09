#ifdef __cplusplus
extern "C" {
#endif

int QRectF_Contains(void* ptr, void* point);
int QRectF_Contains3(void* ptr, void* rectangle);
int QRectF_Intersects(void* ptr, void* rectangle);
void* QRectF_NewQRectF();
void* QRectF_NewQRectF3(void* topLeft, void* bottomRight);
void* QRectF_NewQRectF2(void* topLeft, void* size);
void* QRectF_NewQRectF5(void* rectangle);
void* QRectF_NewQRectF4(double x, double y, double width, double height);
void QRectF_Adjust(void* ptr, double dx1, double dy1, double dx2, double dy2);
double QRectF_Bottom(void* ptr);
int QRectF_Contains2(void* ptr, double x, double y);
double QRectF_Height(void* ptr);
int QRectF_IsEmpty(void* ptr);
int QRectF_IsNull(void* ptr);
int QRectF_IsValid(void* ptr);
double QRectF_Left(void* ptr);
void QRectF_MoveBottom(void* ptr, double y);
void QRectF_MoveBottomLeft(void* ptr, void* position);
void QRectF_MoveBottomRight(void* ptr, void* position);
void QRectF_MoveCenter(void* ptr, void* position);
void QRectF_MoveLeft(void* ptr, double x);
void QRectF_MoveRight(void* ptr, double x);
void QRectF_MoveTo2(void* ptr, void* position);
void QRectF_MoveTo(void* ptr, double x, double y);
void QRectF_MoveTop(void* ptr, double y);
void QRectF_MoveTopLeft(void* ptr, void* position);
void QRectF_MoveTopRight(void* ptr, void* position);
double QRectF_Right(void* ptr);
void QRectF_SetBottom(void* ptr, double y);
void QRectF_SetBottomLeft(void* ptr, void* position);
void QRectF_SetBottomRight(void* ptr, void* position);
void QRectF_SetCoords(void* ptr, double x1, double y1, double x2, double y2);
void QRectF_SetHeight(void* ptr, double height);
void QRectF_SetLeft(void* ptr, double x);
void QRectF_SetRect(void* ptr, double x, double y, double width, double height);
void QRectF_SetRight(void* ptr, double x);
void QRectF_SetSize(void* ptr, void* size);
void QRectF_SetTop(void* ptr, double y);
void QRectF_SetTopLeft(void* ptr, void* position);
void QRectF_SetTopRight(void* ptr, void* position);
void QRectF_SetWidth(void* ptr, double width);
void QRectF_SetX(void* ptr, double x);
void QRectF_SetY(void* ptr, double y);
double QRectF_Top(void* ptr);
void QRectF_Translate2(void* ptr, void* offset);
void QRectF_Translate(void* ptr, double dx, double dy);
double QRectF_Width(void* ptr);
double QRectF_X(void* ptr);
double QRectF_Y(void* ptr);

#ifdef __cplusplus
}
#endif