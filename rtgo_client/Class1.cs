using System;
using System.Collections.Generic;
using System.Linq;
using Autodesk.DesignScript.Interfaces;

namespace rtgo_client
{

    public class RTGOTestClient
    {

        private static List<List<T>> Split<T>(List<T> source, int subListLength)
        {
            return source.
               Select((x, i) => new { Index = i, Value = x })
               .GroupBy(x => x.Index / subListLength)
               .Select(x => x.Select(v => v.Value).ToList())
               .ToList();
        }

        public static List<DSCore.Color> Render(List<Autodesk.DesignScript.Interfaces.IGraphicItem> geos)
        {
            // The port number(5001) must match the port of the gRPC server.
            var channel = new Grpc.Core.Channel("https://localhost:50051", Grpc.Core.ChannelCredentials.Insecure);
            var client = new Rtgo.rtgo.rtgoClient(channel);

            var camera = new Rtgo.Camera();
            camera.Lookfrom = new Rtgo.Vec3() { X = 0, Y = -2, Z = -5 };
            camera.LookAt = new Rtgo.Vec3() { X = 0, Y = 0, Z = 0 };
            camera.Vup = new Rtgo.Vec3() { X = 0, Y = 1, Z = 0 };
            camera.FovDegrees = 45.0;
            camera.AspectRatio = 4.0 / 3.0;
            camera.Aperture = .0000001;
            camera.Focuslength = 4;

            var packages = new List<IRenderPackage>();
            var hittables = new List<Rtgo.Hittable>();
            var materials = new List<Rtgo.Material>();
            var factory = new Dynamo.Visualization.DefaultRenderPackageFactory();

            var tp = new TessellationParameters();
            foreach (var geo in geos)
            {
                var rp = factory.CreateRenderPackage();
                geo.Tessellate(rp, tp);
                packages.Add(rp);
            }

            //grab double components from rp and subset them into points and further into triangles


            foreach (var package in packages)
            {
                var mesh = new Rtgo.Hittable();
                mesh.HittableType = Rtgo.Hittable.Types.HittableType.Mesh;
                var fd = new Rtgo.FaceData();
                fd.Vertindicies.AddRange(package.MeshIndices);
                fd.Normalindicies.AddRange(package.MeshIndices);
                fd.Texcoordindicies.AddRange(package.MeshIndices);

                var vertData = Split(package.MeshVertices.ToList(), 3);
                var vec3s = vertData.Select(x => new Rtgo.Vec3() { X = x[0], Y = x[1], Z = x[2] }).ToList();
                var normData = Split(package.MeshNormals.ToList(), 3);
                var vec3sN = vertData.Select(x => new Rtgo.Vec3() { X = x[0], Y = x[1], Z = x[2] }).ToList();
                var TexData = Split(package.MeshTextureCoordinates.ToList(), 3);
                var vec3sT = vertData.Select(x => new Rtgo.Vec3() { X = x[0], Y = x[1], Z = x[2] }).ToList();

                mesh.Verts.AddRange(vec3s);
                mesh.Texcoords.AddRange(vec3sT);
                mesh.Normals.AddRange(vec3sN);
                hittables.Add(mesh);
                var mat = new Rtgo.Material();
                mat.MaterialType = Rtgo.Material.Types.MaterialType.DiffuseMaterial;
                var tex = new Rtgo.Texture();
                tex.Colors.Add(new Rtgo.Vec3 { X = .4, Y = .7, Z = .4 });
                mat.Texture = tex;
                materials.Add(mat);
            }

            var renderRequest = new Rtgo.RenderRequest();
            renderRequest.Camera = camera;
            renderRequest.Hittable.AddRange(hittables);
            renderRequest.Material.AddRange(materials);


            var reply = client.Render(renderRequest);

            Console.WriteLine($"recieved {reply.Colors.Count} colors");


            return reply.Colors.Select(x=> DSCore.Color.ByARGB(255,(int)x.X,(int)x.Y,(int)x.Z)).ToList();

        }
    }
}
