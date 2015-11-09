#ifdef __cplusplus
extern "C" {
#endif

void* QTextImageFormat_NewQTextImageFormat();
double QTextImageFormat_Height(void* ptr);
int QTextImageFormat_IsValid(void* ptr);
char* QTextImageFormat_Name(void* ptr);
void QTextImageFormat_SetHeight(void* ptr, double height);
void QTextImageFormat_SetName(void* ptr, char* name);
void QTextImageFormat_SetWidth(void* ptr, double width);
double QTextImageFormat_Width(void* ptr);

#ifdef __cplusplus
}
#endif