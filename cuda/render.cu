#include "render.h"
#include <cuda_runtime.h>
#include <stdio.h>
__global__
void renderPixel(int n,int *rout, int *gout, int *bout){
    int index = threadIdx.x;
    int stride = blockDim.x;
    for (int i = index; i < n; i += stride)
    {
        rout[i] = index;
        gout[i] = index;
        bout[i] = index;
    }
}



void wrapper(int width, int height, int *xs, int *ys, int *rout, int *gout, int *bout){
    
    int N = width * height;
    // Allocate Unified Memory – accessible from CPU or GPU
int *routGPU, *goutGPU, *boutGPU;

  cudaMallocManaged(&routGPU, N*sizeof(int));
  cudaMallocManaged(&goutGPU, N*sizeof(int));
  cudaMallocManaged(&boutGPU, N*sizeof(int));

   // initialize x and y arrays on the host
   for (int i = 0; i < N; i++) {
    routGPU[i] = 0;
    goutGPU[i] = 0;
    boutGPU[i] = 0;

  }
  int blockSize = 256;
  int numBlocks = (N + blockSize - 1) / blockSize;

  renderPixel<<<numBlocks,blockSize>>>(N,routGPU,goutGPU,boutGPU);

  // Wait for GPU to finish before accessing on host
  cudaDeviceSynchronize();

  //copy memory to go
cudaMemcpy(rout,routGPU,N*sizeof(int),cudaMemcpyDeviceToHost);
cudaMemcpy(gout,goutGPU,N*sizeof(int),cudaMemcpyDeviceToHost);
cudaMemcpy(bout,boutGPU,N*sizeof(int),cudaMemcpyDeviceToHost);


   // Free memory
   cudaFree(routGPU);
   cudaFree(boutGPU);
   cudaFree(goutGPU);
    
}