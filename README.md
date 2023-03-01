# golang-csharp-gRPC-kubectl-v1

## Project 1
c# gRPC client <-> server and client <-> server
1. use "GrpcChannel.ForAddress("http://secondgrpcservice-project1:50102")" to create a lower service channel within default created file
2. In default file, create client object and call its rpc functions
3. ------------
4. Create QHDataService.cs and use it to call Mockqhd.QHDmocker.QHDmockerClient rpc function
5. Switch to use AddScroped to add QHDataService and use AddGrpcClient to add Mockqhd.QHDmocker.QHDmockerClient in program.cs
6. Mockqhd is the namespace that matches to package name in proto file, QHDmocker matches to service name
7. kubectl create -f secondgRPCService-project2.yaml
8. kubectl create -f firstgrpcservice-project1.yaml
9. kubectl create -f firstgrpcclient-project1.yaml

## Project 2
golang client <- gRPC -> c# server

## Useful Sources
1. https://www.infoq.com/articles/getting-started-grpc-dotnet/
2. https://swaminathanvetri.in/2020/01/29/getting-started-with-grpc-dotnet-in-vs-code/
3. https://www.codemag.com/Article/2212071/A-Deep-Dive-into-Working-with-gRPC-in-.NET-6
4. https://www.dotnetcurry.com/aspnet-core/kubernetes-for-developers
5. https://learn.microsoft.com/en-us/aspnet/core/tutorials/grpc/grpc-start?view=aspnetcore-6.0&tabs=visual-studio-code
6. https://learn.microsoft.com/en-us/dotnet/core/docker/build-container?tabs=windows
7. https://learn.microsoft.com/en-us/aspnet/core/fundamentals/minimal-apis?view=aspnetcore-6.0
8. https://medium.com/aspnetrun/using-grpc-in-microservices-for-building-a-high-performance-interservice-communication-with-net-5-11f3e5fa0e9d