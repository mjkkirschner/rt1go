A raytracer implemeted in go - for learning.
References:
[_Ray Tracing in One Weekend_](https://raytracing.github.io/books/RayTracingInOneWeekend.html)

### TODO:
- [x] send rays towards light objects
- [x] specular/refecltive rays for metal etc.
- [x] multicore - samples are split into goroutines, but other strategies are untested
- [x] positionable camera
- [x] BVH or other spatial data structure
- [] specular highlights
- [] update view progress
- [x] ray tri intersection
- [~] obj/stl loading
- [x] texture coords / textures
- [] signed distance field to mesh gen using https://github.com/deadsy/sdfx

![an image](./test.png)
