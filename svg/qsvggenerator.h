#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QSvgGenerator_Description(QtObjectPtr ptr);
char* QSvgGenerator_FileName(QtObjectPtr ptr);
QtObjectPtr QSvgGenerator_OutputDevice(QtObjectPtr ptr);
int QSvgGenerator_Resolution(QtObjectPtr ptr);
void QSvgGenerator_SetDescription(QtObjectPtr ptr, char* description);
void QSvgGenerator_SetFileName(QtObjectPtr ptr, char* fileName);
void QSvgGenerator_SetOutputDevice(QtObjectPtr ptr, QtObjectPtr outputDevice);
void QSvgGenerator_SetResolution(QtObjectPtr ptr, int dpi);
void QSvgGenerator_SetSize(QtObjectPtr ptr, QtObjectPtr size);
void QSvgGenerator_SetTitle(QtObjectPtr ptr, char* title);
void QSvgGenerator_SetViewBox(QtObjectPtr ptr, QtObjectPtr viewBox);
void QSvgGenerator_SetViewBox2(QtObjectPtr ptr, QtObjectPtr viewBox);
char* QSvgGenerator_Title(QtObjectPtr ptr);
QtObjectPtr QSvgGenerator_NewQSvgGenerator();
void QSvgGenerator_DestroyQSvgGenerator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif