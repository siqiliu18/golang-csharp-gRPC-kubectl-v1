using Grpc.Core;
using QhDataService; // This matches qh_data_grpc_service.proto "package qh_data_service;"
using Grpc.Net.Client;
using TemplateServer;

namespace TemplateServer.Services;

public class TemplateService : TemplateMaker.TemplateMakerBase
{
    private readonly ILogger<TemplateService> _logger;
    private QHData.QHDataClient _client;

    public TemplateService(ILogger<TemplateService> logger)
    {
        _logger = logger;

        var channel = GrpcChannel.ForAddress("http://mockqhd-pptservice:50103");
        _client = new QHData.QHDataClient(channel);
    }

    public override Task<TemplateResponse> RetrieveTemplate(TemplateRequest request, ServerCallContext context)
    {
        Console.WriteLine(" >>> In RetrieveTemplate function");

        var req = new QILRequest
        {
            QilID = 6,
            QilVersion = 1,
        };

        Console.WriteLine(" Call QIL before: ");
        var res = _client.CallQil(req);
        Console.WriteLine(res.Data);
        Console.WriteLine(" QilCall works? " + res.Data);

        return Task.FromResult(new TemplateResponse
        {
            Message = "Hello " + request.Name
        });
    }
}
