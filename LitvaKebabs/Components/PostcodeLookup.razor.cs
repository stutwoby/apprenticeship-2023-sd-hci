using LitvaKebabs.Models;
using Microsoft.AspNetCore.Components;
using Microsoft.Extensions.Logging;
using Microsoft.JSInterop;
using Newtonsoft.Json;
using RestSharp;
using System.Web;

namespace LitvaKebabs.Components
{
    public partial class PostcodeLookup
    {
        private const string deliveringFromPostcode = "BH243FW";
        [SupplyParameterFromQuery(Name = "q")]
        public string DeliveringToPostcode { get; set; } = string.Empty;

        protected override async Task OnInitializedAsync()
        {
            // So we firstly need to validate the postcode. Maybe we could utilise the API to do this and be lazy about it?
            // Ok lets do that...
            string deliveringToPostcode = DeliveringToPostcode;
            var postcodeInformation = PostcodeAPIRequest($"api.postcodes.io/postcodes/{deliveringToPostcode}").GetAwaiter().GetResult();
            PostcodeModel.Root postcodeModel = JsonConvert.DeserializeObject<PostcodeModel.Root>(postcodeInformation.ToString());
            if(postcodeModel != null )
            {
                if(postcodeModel.status != 200)
                {
                    string not200Error = $"Error, API returned status of {postcodeModel.status}: " +
                        $"failed to get API contents, wrong postcode potentially entered. " +
                        $"Please reload the page and try again.";
                    await jsRuntime.InvokeVoidAsync("alert", not200Error);
                    logger.LogError(not200Error);
                    throw new Exception();
                }
            }
            else
            {
                throw new Exception("Error: postcodeModel is null.");
            }
            CalculateDistance();
        }

        private async Task<RestResponse?> PostcodeAPIRequest(string apiUri)
        {
            var options = new RestClientOptions(HttpUtility.UrlEncode(apiUri));
            var client = new RestClient(options);
            var request = new RestRequest("");
            request.AddHeader("accept", "application/json");
            var response = await client.PostAsync(request);
            return response;
        }
        private decimal CalculateDistance()
        {
            return 0;
        }
    }
}
