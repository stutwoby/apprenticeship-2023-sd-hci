using LitvaKebabs.Models;
using Microsoft.AspNetCore.Components;
using Microsoft.JSInterop;
using Newtonsoft.Json;
using RestSharp;
using System.Web;

namespace LitvaKebabs.Components
{
    public partial class PostcodeLookup
    {
        private const string deliveringFromPostcode = "BH243FW";
        private const double deliveringFromPostcodeLon = -1.775614;
        private const double deliveringFromPostcodeLat = 50.84143; 
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
                else
                {

                }
            }
            else
            {
                throw new Exception("Error: postcodeModel is null.");
            }
            
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
        // From https://stackoverflow.com/a/51839058
        // This function is covered by CC-BY-SA 4.0
        // Modified to return distance in miles.
        public double GetDistance(double longitude, double latitude, double otherLongitude, double otherLatitude)
        {
            longitude = deliveringFromPostcodeLon;
            latitude = deliveringFromPostcodeLat;
            var d1 = latitude * (Math.PI / 180.0);
            var num1 = longitude * (Math.PI / 180.0);
            var d2 = otherLatitude * (Math.PI / 180.0);
            var num2 = otherLongitude * (Math.PI / 180.0) - num1;
            var d3 = Math.Pow(Math.Sin((d2 - d1) / 2.0), 2.0) + Math.Cos(d1) * Math.Cos(d2) * Math.Pow(Math.Sin(num2 / 2.0), 2.0);
            var res = 6376500.0 * (2.0 * Math.Atan2(Math.Sqrt(d3), Math.Sqrt(1.0 - d3)));
            return res * 0.00062137;
        }
    }
}
