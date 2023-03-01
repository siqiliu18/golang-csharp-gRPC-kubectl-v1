namespace TemplateServer.Services
{
    public class MockDataService
    {
        private readonly ILogger<MockDataService> _logger;
        private QhDataService.QHData.QHDataClient _client;

        public MockDataService(ILogger<MockDataService> logger, QhDataService.QHData.QHDataClient client)
        {
            _logger = logger;
            _client = client;
        }

        public string QilCall()
        {
            Console.WriteLine(">>> QilCall");
            var req = new QhDataService.QILRequest
            {
                QilID = 6,
                QilVersion = 1,
            };
            Console.WriteLine("!!! actual rpc call !!!");
            var res = _client.CallQil(req).Data;
            Console.WriteLine("@@@ res = " + res);
            return res;
        }
    }
}