#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMimeData_NewQMimeData();
void QMimeData_Clear(QtObjectPtr ptr);
char* QMimeData_ColorData(QtObjectPtr ptr);
char* QMimeData_Formats(QtObjectPtr ptr);
int QMimeData_HasColor(QtObjectPtr ptr);
int QMimeData_HasFormat(QtObjectPtr ptr, char* mimeType);
int QMimeData_HasHtml(QtObjectPtr ptr);
int QMimeData_HasImage(QtObjectPtr ptr);
int QMimeData_HasText(QtObjectPtr ptr);
int QMimeData_HasUrls(QtObjectPtr ptr);
char* QMimeData_Html(QtObjectPtr ptr);
char* QMimeData_ImageData(QtObjectPtr ptr);
void QMimeData_RemoveFormat(QtObjectPtr ptr, char* mimeType);
void QMimeData_SetColorData(QtObjectPtr ptr, char* color);
void QMimeData_SetData(QtObjectPtr ptr, char* mimeType, QtObjectPtr data);
void QMimeData_SetHtml(QtObjectPtr ptr, char* html);
void QMimeData_SetImageData(QtObjectPtr ptr, char* image);
void QMimeData_SetText(QtObjectPtr ptr, char* text);
char* QMimeData_Text(QtObjectPtr ptr);
void QMimeData_DestroyQMimeData(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif