#ifdef __cplusplus
extern "C" {
#endif

void* QTextFrame_NewQTextFrame(void* document);
int QTextFrame_FirstPosition(void* ptr);
int QTextFrame_LastPosition(void* ptr);
void* QTextFrame_ParentFrame(void* ptr);
void QTextFrame_SetFrameFormat(void* ptr, void* format);
void QTextFrame_DestroyQTextFrame(void* ptr);

#ifdef __cplusplus
}
#endif