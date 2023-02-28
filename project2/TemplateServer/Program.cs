using Microsoft.AspNetCore.Server.Kestrel.Core;
using TemplateServer.Services;

var builder = WebApplication.CreateBuilder(args);

// Additional configuration is required to successfully run gRPC on macOS.
// For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682
builder.WebHost.ConfigureKestrel(options =>
{
    // Setup a HTTP/2 endpoint without TLS.
    // !!! *** https://github.com/grpc/grpc-dotnet/issues/1874 *** !!!
    options.ListenAnyIP(5288, o => o.Protocols = HttpProtocols.Http2);
});

// Add services to the container.
builder.Services.AddGrpc();

// builder.Services.AddGrpcClient<QHDataService.QHData.QHDataClient>(o =>
// {
//     o.Address = new Uri("http://mockqhd-pptservice:7003");
// });

var app = builder.Build();

// Configure the HTTP request pipeline.
app.MapGrpcService<TemplateService>();
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");

app.Run();
