using LitvaKebabs.Pages;
using Microsoft.AspNetCore.Components;

namespace LitvaKebabs.Components
{
    public partial class OrderButton
    {
        [Parameter]
        public Dictionary<string, decimal> tableData { get; set; } = new Dictionary<string, decimal>();
        OrderMenuRoot orderMenuRoot = new OrderMenuRoot();
        private void AddMenuItemToReciept()
        {
            foreach(var item in orderMenuRoot.MenuItems)
            {
                tableData.Add(item.Key, item.Value);
                StateHasChanged();
            }
            
        }
    }
}
