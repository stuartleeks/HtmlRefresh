FROM microsoft/dotnet:2.0.0-sdk as build-ENV
WORKDIR /app

COPY /HtmlRefresh/HtmlRefresh.csproj ./
RUN dotnet restore --runtime linux-x64

COPY /HtmlRefresh/ ./
RUN dotnet build --runtime linux-x64
RUN dotnet publish -c Release -o out --no-restore --runtime linux-x64


# runtime image:
FROM microsoft/dotnet:2.0.0-runtime-deps
WORKDIR /app
COPY --from=build-env /app/out ./

EXPOSE 5000/tcp

ARG ARG_HTML_REFRESH_BGCOLOUR
ARG ARG_HTML_REFRESH_FGCOLOUR

ENV HTML_REFRESH_BGCOLOUR $ARG_HTML_REFRESH_BGCOLOUR
ENV HTML_REFRESH_FGCOLOUR $ARG_HTML_REFRESH_FGCOLOUR
ENV HTML_REFRESH_MESSAGE "%HOSTNAME%"
ENV ASPNETCORE_URLS=http://*:5000

ENTRYPOINT ["./HtmlRefresh"]

