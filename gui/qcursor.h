#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QCursor_QCursor_SetPos2(QtObjectPtr screen, int x, int y);
void QCursor_QCursor_SetPos(int x, int y);
QtObjectPtr QCursor_NewQCursor();
QtObjectPtr QCursor_NewQCursor6(QtObjectPtr other);
QtObjectPtr QCursor_NewQCursor2(int shape);
QtObjectPtr QCursor_NewQCursor3(QtObjectPtr bitmap, QtObjectPtr mask, int hotX, int hotY);
QtObjectPtr QCursor_NewQCursor5(QtObjectPtr c);
QtObjectPtr QCursor_NewQCursor4(QtObjectPtr pixmap, int hotX, int hotY);
QtObjectPtr QCursor_Bitmap(QtObjectPtr ptr);
QtObjectPtr QCursor_Mask(QtObjectPtr ptr);
void QCursor_QCursor_SetPos4(QtObjectPtr screen, QtObjectPtr p);
void QCursor_QCursor_SetPos3(QtObjectPtr p);
void QCursor_SetShape(QtObjectPtr ptr, int shape);
int QCursor_Shape(QtObjectPtr ptr);
void QCursor_DestroyQCursor(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif