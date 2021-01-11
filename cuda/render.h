
#ifdef __cplusplus
extern "C" {  // only need to export C interface if
              // used by C++ source code
#endif

 __declspec(dllexport) void wrapper(int width, int height, int *xs, int *ys, int *rout, int *gout, int *bout);
 
#ifdef __cplusplus
}
#endif