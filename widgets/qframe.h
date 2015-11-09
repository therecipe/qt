#ifdef __cplusplus
extern "C" {
#endif

int QFrame_FrameShadow(void* ptr);
int QFrame_FrameShape(void* ptr);
int QFrame_FrameWidth(void* ptr);
int QFrame_LineWidth(void* ptr);
int QFrame_MidLineWidth(void* ptr);
void QFrame_SetFrameRect(void* ptr, void* v);
void QFrame_SetFrameShadow(void* ptr, int v);
void QFrame_SetFrameShape(void* ptr, int v);
void QFrame_SetLineWidth(void* ptr, int v);
void QFrame_SetMidLineWidth(void* ptr, int v);
void* QFrame_NewQFrame(void* parent, int f);
int QFrame_FrameStyle(void* ptr);
void QFrame_SetFrameStyle(void* ptr, int style);
void QFrame_DestroyQFrame(void* ptr);

#ifdef __cplusplus
}
#endif