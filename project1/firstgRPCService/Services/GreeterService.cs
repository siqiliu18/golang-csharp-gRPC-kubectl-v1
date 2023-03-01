using Grpc.Core;
using firstgRPCService;
using Grpc.Net.Client;
using Mockqhd;

namespace firstgRPCService.Services;

public class GreeterService : Greeter.GreeterBase
{
    private readonly ILogger<GreeterService> _logger;
    // private QHDmocker.QHDmockerClient _client;
    private QHDataService _dataService;
    public GreeterService(ILogger<GreeterService> logger, QHDataService dataService)
    {
        _logger = logger;
        // var channel = GrpcChannel.ForAddress("http://secondgrpcservice-project1:50102");
        // _client = new QHDmocker.QHDmockerClient(channel);

        _dataService = dataService;
    }

    public override Task<HelloReply> SayHello(HelloRequest request, ServerCallContext context)
    {
        var req = new QilRequest
        {
            QilID = 6,
            QilVersion = 1,
        };

        Console.WriteLine("? Call QHDataServie API before: ");
        string str = _dataService.CallQILnothing();
        Console.WriteLine(str);
        Console.WriteLine("? Call QHDataService API after. ");

        // Console.WriteLine("Call QIL Before: ");
        // var res = _client.CallQIL(req);
        // Console.WriteLine(res.Data);
        // Console.WriteLine("Call QIL After.");

        return Task.FromResult(new HelloReply
        {
            Message = "Hello " + request.Name
        });
    }
}
