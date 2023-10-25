using LitvaKebabs.Models;
using Microsoft.AspNetCore.Components;
using Newtonsoft.Json;

namespace LitvaKebabs.Pages
{
    public partial class OrderMenuRoot
    {
        public OrderMenuRoot()
        {
            MenuItems = _menuItems();
            SauceItems = _sauceItems();
        }

        [Parameter]
        public Dictionary<string, decimal> MenuItems { get; set; }
        [Parameter]
        public Dictionary<string, decimal> SauceItems { get; set; }

        private Dictionary<string, decimal> _menuItems()
        {
            string json = File.ReadAllText(@$"{Environment.CurrentDirectory}\KebabMenu.json");
            List<KebabMenuModel> data = JsonConvert.DeserializeObject<List<KebabMenuModel>>(json);

            Dictionary<string, decimal> menuItems = data.ToDictionary(keySelector: data => data.Item, elementSelector: records => records.Price);
            return menuItems;
        }

        private Dictionary<string, decimal> _sauceItems()
        {
            string json = File.ReadAllText(@$"{Environment.CurrentDirectory}\SauceMenu.json");
            List<SauceMenuModel> data = JsonConvert.DeserializeObject<List<SauceMenuModel>>(json);
            Dictionary<string, decimal> sauceItems = data.ToDictionary(keySelector: records => records.Sauce, elementSelector: records => records.Price);
            return sauceItems;
        }
    }
}
