# golang-csharp-gRPC-kubectl-v1

## Project 1
c# gRPC client and server

## Project 2
golang client <- gRPC -> c# server
1. The 2 proto files have to be the same!
2. create docker file for csharp server: docker build -t secondgrpcservice-project2-multi -f Dockerfile .
3. create docker file for mock qhd server: docker build -t pptp2-modkqhd-image -f Dockerfile .
4. launch services and pods
5. use: var channel = GrpcChannel.ForAddress("http://mockqhd-pptservice:50103"); 
6. assign to private _client = new QHData.QHDataClient(channel); that is belonged to default file
7. call RPC function of mock qhd

## Useful Sources
1. https://www.infoq.com/articles/getting-started-grpc-dotnet/
2. https://swaminathanvetri.in/2020/01/29/getting-started-with-grpc-dotnet-in-vs-code/
3. https://www.codemag.com/Article/2212071/A-Deep-Dive-into-Working-with-gRPC-in-.NET-6
4. https://www.dotnetcurry.com/aspnet-core/kubernetes-for-developers
5. https://learn.microsoft.com/en-us/aspnet/core/tutorials/grpc/grpc-start?view=aspnetcore-6.0&tabs=visual-studio-code
6. https://learn.microsoft.com/en-us/dotnet/core/docker/build-container?tabs=windows
7. https://learn.microsoft.com/en-us/aspnet/core/fundamentals/minimal-apis?view=aspnetcore-6.0
8. https://medium.com/aspnetrun/using-grpc-in-microservices-for-building-a-high-performance-interservice-communication-with-net-5-11f3e5fa0e9d

## Setup workspace for multiple modules in one repo
1. https://github.com/golang/tools/blob/master/gopls/doc/workspace.md
2. https://github.com/golang/vscode-go/issues/2132