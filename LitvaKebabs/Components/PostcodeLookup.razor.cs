using Microsoft.AspNetCore.Components;

namespace LitvaKebabs.Components
{
    public partial class PostcodeLookup
    {
        private const string deliveringFromPostcode = "BH243FW";
        [SupplyParameterFromQuery(Name="q")]
        public string SearchTerm { get; set; }

        protected override async Task OnInitializedAsync()
        {

        }
        private decimal CalculateDistance()
        {
            return 0;
        }
    }
}
