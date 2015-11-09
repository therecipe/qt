#ifdef __cplusplus
extern "C" {
#endif

void* QTextOption_NewQTextOption3(void* other);
void* QTextOption_NewQTextOption();
void* QTextOption_NewQTextOption2(int alignment);
int QTextOption_Alignment(void* ptr);
int QTextOption_Flags(void* ptr);
void QTextOption_SetAlignment(void* ptr, int alignment);
void QTextOption_SetFlags(void* ptr, int flags);
void QTextOption_SetTabStop(void* ptr, double tabStop);
void QTextOption_SetTextDirection(void* ptr, int direction);
void QTextOption_SetUseDesignMetrics(void* ptr, int enable);
void QTextOption_SetWrapMode(void* ptr, int mode);
double QTextOption_TabStop(void* ptr);
int QTextOption_TextDirection(void* ptr);
int QTextOption_UseDesignMetrics(void* ptr);
int QTextOption_WrapMode(void* ptr);
void QTextOption_DestroyQTextOption(void* ptr);

#ifdef __cplusplus
}
#endif