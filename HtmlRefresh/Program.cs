using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Http;
using System;
using System.Threading.Tasks;

namespace HtmlRefresh
{
    public class Program
    {
        private static string Html;

        public static void Main(string[] args)
        {
            Console.WriteLine("HtmlRefresh Starting...");
            BuildHtmlFromOptions();

            Console.WriteLine("Start pipeline.");

            var host = new WebHostBuilder()
                .UseKestrel()
                .Configure(app => app.Run(RenderHtml))
                .Build();

            host.Run();
        }

        private static void BuildHtmlFromOptions()
        {
            Console.WriteLine("Options: ");

            var backgroundColour = Environment.GetEnvironmentVariable("HTML_REFRESH_BGCOLOUR") ?? "azure";
            Console.WriteLine($"\tBgColour: {backgroundColour}");

            var foregroundColour = Environment.GetEnvironmentVariable("HTML_REFRESH_FGCOLOUR") ?? "black";
            Console.WriteLine($"\tFgColour: {foregroundColour}");

            var fontSize = Environment.GetEnvironmentVariable("HTML_REFRESH_FONTSIZE") ?? "16vw";
            Console.WriteLine($"\tFontSize: {fontSize}");

            var refreshRate = Environment.GetEnvironmentVariable("HTML_REFRESH_RATE") ?? "1";
            Console.WriteLine($"\tRate: {refreshRate}");

            var message = Environment.GetEnvironmentVariable("HTML_REFRESH_MESSAGE") ?? "$HOSTNAME";
            var messageExpanded = Environment.ExpandEnvironmentVariables(message);
            Console.WriteLine($"\tMessage: '{message}' (expanded: '{messageExpanded}')");

            Html = $"<html><head><meta http-equiv=\"refresh\" content=\"{refreshRate}\"></head><body style=\"background-color: {backgroundColour};color: {foregroundColour};font-size: {fontSize};font-family: 'Open Sans',sans-serif; margin:0;padding:20px\">{messageExpanded}</body></head>";
        }

        private static async Task RenderHtml(HttpContext context)
        {
            await context.Response.WriteAsync(Html);
        }

    }
}
