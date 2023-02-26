// See https://aka.ms/new-console-template for more information
using firstgRPCClient;
using Grpc.Net.Client;

var channel = GrpcChannel.ForAddress("http://firstgrpcservice-project1:50101");

var client = new Greeter.GreeterClient(channel);

Console.WriteLine("Hello, World!");

var response = await client.SayHelloAsync(new HelloRequest
{
    Name = "Siqi"
});

Console.WriteLine("From Server: " + response.Message);
