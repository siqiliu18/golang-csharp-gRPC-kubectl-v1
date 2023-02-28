using Grpc.Core;
using secondgRPCService;

namespace secondgRPCService.Services;

public class GreeterService : QHDmocker.QHDmockerBase
{
    private readonly ILogger<GreeterService> _logger;
    public GreeterService(ILogger<GreeterService> logger)
    {
        _logger = logger;
    }

    public override Task<QilReply> CallQIL(QilRequest request, ServerCallContext context)
    {
        return Task.FromResult(new QilReply
        {
            Data = "QIL " + request.QilID.ToString() + "v" + request.QilVersion.ToString(),
        });
    }
}
