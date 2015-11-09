#ifdef __cplusplus
extern "C" {
#endif

void QCursor_QCursor_SetPos2(void* screen, int x, int y);
void QCursor_QCursor_SetPos(int x, int y);
void* QCursor_NewQCursor();
void* QCursor_NewQCursor6(void* other);
void* QCursor_NewQCursor2(int shape);
void* QCursor_NewQCursor3(void* bitmap, void* mask, int hotX, int hotY);
void* QCursor_NewQCursor5(void* c);
void* QCursor_NewQCursor4(void* pixmap, int hotX, int hotY);
void* QCursor_Bitmap(void* ptr);
void* QCursor_Mask(void* ptr);
void QCursor_QCursor_SetPos4(void* screen, void* p);
void QCursor_QCursor_SetPos3(void* p);
void QCursor_SetShape(void* ptr, int shape);
int QCursor_Shape(void* ptr);
void QCursor_DestroyQCursor(void* ptr);

#ifdef __cplusplus
}
#endif