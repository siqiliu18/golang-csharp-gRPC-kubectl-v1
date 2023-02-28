using Grpc.Core;
using firstgRPCService;
using Grpc.Net.Client;
using Mockqhd;

namespace firstgRPCService.Services;

public class GreeterService : Greeter.GreeterBase
{
    private readonly ILogger<GreeterService> _logger;
    private QHDmocker.QHDmockerClient _client;
    public GreeterService(ILogger<GreeterService> logger)
    {
        _logger = logger;
        var channel = GrpcChannel.ForAddress("secondgrpcservice-project1:50102");
        _client = new QHDmocker.QHDmockerClient(channel);
    }

    public override Task<HelloReply> SayHello(HelloRequest request, ServerCallContext context)
    {
        var req = new QilRequest
        {
            QilID = 6,
            QilVersion = 1,
        };

        Console.WriteLine("Before: ");
        var res = _client.CallQIL(req);
        Console.WriteLine(res.Data);
        Console.WriteLine("After.");

        return Task.FromResult(new HelloReply
        {
            Message = "Hello " + request.Name
        });
    }
}
