#ifdef __cplusplus
extern "C" {
#endif

void* QMimeData_NewQMimeData();
void QMimeData_Clear(void* ptr);
void* QMimeData_ColorData(void* ptr);
void* QMimeData_Data(void* ptr, char* mimeType);
char* QMimeData_Formats(void* ptr);
int QMimeData_HasColor(void* ptr);
int QMimeData_HasFormat(void* ptr, char* mimeType);
int QMimeData_HasHtml(void* ptr);
int QMimeData_HasImage(void* ptr);
int QMimeData_HasText(void* ptr);
int QMimeData_HasUrls(void* ptr);
char* QMimeData_Html(void* ptr);
void* QMimeData_ImageData(void* ptr);
void QMimeData_RemoveFormat(void* ptr, char* mimeType);
void QMimeData_SetColorData(void* ptr, void* color);
void QMimeData_SetData(void* ptr, char* mimeType, void* data);
void QMimeData_SetHtml(void* ptr, char* html);
void QMimeData_SetImageData(void* ptr, void* image);
void QMimeData_SetText(void* ptr, char* text);
char* QMimeData_Text(void* ptr);
void QMimeData_DestroyQMimeData(void* ptr);

#ifdef __cplusplus
}
#endif