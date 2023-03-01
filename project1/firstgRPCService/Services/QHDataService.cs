using Mockqhd;

public class QHDataService
{
    private readonly ILogger<QHDataService> _logger;
    private QHDmocker.QHDmockerClient _client;
    
    public QHDataService(ILogger<QHDataService> logger, QHDmocker.QHDmockerClient client)
    {
        _logger = logger;
        _client = client;
    }

    public string CallQILnothing()
    {
        Console.WriteLine("+++CallQI+++");
        var req = new Mockqhd.QilRequest
        {
            QilID = 6,
            QilVersion = 1,
        };

        Console.WriteLine("!!! actually rpc call !!!");
        var res = _client.CallQIL(req).Data;
        Console.WriteLine("@@@ res = ", res);

        return res;
    }
}