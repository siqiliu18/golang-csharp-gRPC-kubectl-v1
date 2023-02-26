using Grpc.Core;
using TemplateServer;

namespace TemplateServer.Services;

public class TemplateService : TemplateMaker.TemplateMakerBase
{
    private readonly ILogger<TemplateService> _logger;
    public TemplateService(ILogger<TemplateService> logger)
    {
        _logger = logger;
    }

    public override Task<TemplateResponse> RetrieveTemplate(TemplateRequest request, ServerCallContext context)
    {
        Console.WriteLine(" >>> In RetrieveTemplate function");
        return Task.FromResult(new TemplateResponse
        {
            Message = "Hello " + request.Name
        });
    }
}
