#ifdef __cplusplus
extern "C" {
#endif

void* QPdfWriter_NewQPdfWriter2(void* device);
void* QPdfWriter_NewQPdfWriter(char* filename);
char* QPdfWriter_Creator(void* ptr);
int QPdfWriter_NewPage(void* ptr);
int QPdfWriter_Resolution(void* ptr);
void QPdfWriter_SetCreator(void* ptr, char* creator);
int QPdfWriter_SetPageLayout(void* ptr, void* newPageLayout);
int QPdfWriter_SetPageMargins(void* ptr, void* margins);
int QPdfWriter_SetPageMargins2(void* ptr, void* margins, int units);
int QPdfWriter_SetPageOrientation(void* ptr, int orientation);
int QPdfWriter_SetPageSize(void* ptr, void* pageSize);
void QPdfWriter_SetResolution(void* ptr, int resolution);
void QPdfWriter_SetTitle(void* ptr, char* title);
char* QPdfWriter_Title(void* ptr);
void QPdfWriter_DestroyQPdfWriter(void* ptr);

#ifdef __cplusplus
}
#endif