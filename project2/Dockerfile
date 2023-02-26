FROM mcr.microsoft.com/dotnet/sdk:6.0 AS builder
WORKDIR /src
ADD TemplateServer .
RUN dotnet build TemplateServer.csproj -c Release -o out

FROM mcr.microsoft.com/dotnet/aspnet:6.0
WORKDIR /app
COPY --from=builder /src/out .
ENTRYPOINT ["dotnet", "TemplateServer.dll"]