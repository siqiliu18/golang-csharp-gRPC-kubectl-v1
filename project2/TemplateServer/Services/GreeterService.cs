using Grpc.Core;
using Grpc.Net.Client;
using QHDataService;
using TemplateServer;

namespace TemplateServer.Services;

public class TemplateService : TemplateMaker.TemplateMakerBase
{
    private readonly ILogger<TemplateService> _logger;
    // private QHDataService _dataService;
    private QHData.QHDataClient client;

    public TemplateService(ILogger<TemplateService> logger)
    {
        _logger = logger;
        // _dataService = dataService;

        var channel = GrpcChannel.ForAddress("http://mockqhd-pptservice:7003");
        client = new QHData.QHDataClient(channel);
    }

    public override async Task<TemplateResponse> RetrieveTemplate(TemplateRequest request, ServerCallContext context)
    {
        Console.WriteLine(" >>> In RetrieveTemplate function");

        var req = new QILRequest
        {
            QilID = 6,
            QilVersion = 1,
        };
        var response = await client.CallQILAsync(req);

        // string qilRes = _dataService.QilCall(6, 1);

        Console.WriteLine(" QilCall works? ", response.Data);

        return (new TemplateResponse
        {
            Message = "Hello " + request.Name
        });
    }
}
