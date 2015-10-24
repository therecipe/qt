#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPdfWriter_NewQPdfWriter2(QtObjectPtr device);
QtObjectPtr QPdfWriter_NewQPdfWriter(char* filename);
char* QPdfWriter_Creator(QtObjectPtr ptr);
int QPdfWriter_NewPage(QtObjectPtr ptr);
int QPdfWriter_Resolution(QtObjectPtr ptr);
void QPdfWriter_SetCreator(QtObjectPtr ptr, char* creator);
int QPdfWriter_SetPageLayout(QtObjectPtr ptr, QtObjectPtr newPageLayout);
int QPdfWriter_SetPageMargins(QtObjectPtr ptr, QtObjectPtr margins);
int QPdfWriter_SetPageMargins2(QtObjectPtr ptr, QtObjectPtr margins, int units);
int QPdfWriter_SetPageOrientation(QtObjectPtr ptr, int orientation);
int QPdfWriter_SetPageSize(QtObjectPtr ptr, QtObjectPtr pageSize);
void QPdfWriter_SetResolution(QtObjectPtr ptr, int resolution);
void QPdfWriter_SetTitle(QtObjectPtr ptr, char* title);
char* QPdfWriter_Title(QtObjectPtr ptr);
void QPdfWriter_DestroyQPdfWriter(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif