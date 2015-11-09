#ifdef __cplusplus
extern "C" {
#endif

double QPressureReading_Pressure(void* ptr);
double QPressureReading_Temperature(void* ptr);
void QPressureReading_SetPressure(void* ptr, double pressure);
void QPressureReading_SetTemperature(void* ptr, double temperature);

#ifdef __cplusplus
}
#endif