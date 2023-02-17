using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Threading.Tasks;
using Kurdi.ECommerce.Inventory.Api.Helpers;

namespace Kurdi.ECommerce.Inventory.Api.Middleware
{

    public class LanguageMidleware
    {
        private readonly RequestDelegate _next;
        public LanguageMidleware(RequestDelegate next)
        {
            _next = next;
        }
        public async Task InvokeAsync(HttpContext context)
        {
            List<string> SupportedLanguages = new List<string>() { "ar", "en" };

            bool validLanguage = true;
            string? languageHeader = context.Request.Headers["Language"];
            if (!string.IsNullOrEmpty(languageHeader))
            {
                if (!SupportedLanguages.Contains(languageHeader))
                {
                    context.Response.StatusCode = 404;
                    validLanguage = false;
                    var responseBody = new {
                        successed = false,
                        message = Translator.Translate("VALIDATION:NOT_VALID_LANGUAGE")
                    };
                    await context.Response.WriteAsJsonAsync(responseBody);
                }

                LanguageInfoHelper.CurrentLanguage = languageHeader;
            }
            else
            {
                LanguageInfoHelper.CurrentLanguage = "ar";
            }

            if (validLanguage)
            {
                await _next(context);
            }
        }
    }

    public static class LanguageMidlewareExtensions
    {
        public static IApplicationBuilder UseLanguageMiddleware(
            this IApplicationBuilder builder)
        {
            return builder.UseMiddleware<LanguageMidleware>();
        }
    }
}