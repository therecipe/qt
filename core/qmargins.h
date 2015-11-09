#ifdef __cplusplus
extern "C" {
#endif

void* QMargins_NewQMargins();
void* QMargins_NewQMargins2(int left, int top, int right, int bottom);
int QMargins_Bottom(void* ptr);
int QMargins_IsNull(void* ptr);
int QMargins_Left(void* ptr);
int QMargins_Right(void* ptr);
void QMargins_SetBottom(void* ptr, int bottom);
void QMargins_SetLeft(void* ptr, int left);
void QMargins_SetRight(void* ptr, int right);
void QMargins_SetTop(void* ptr, int Top);
int QMargins_Top(void* ptr);

#ifdef __cplusplus
}
#endif