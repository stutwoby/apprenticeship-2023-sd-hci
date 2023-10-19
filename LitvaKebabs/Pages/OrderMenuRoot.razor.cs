using CsvHelper;
using CsvHelper.Configuration;
using LitvaKebabs.Models;
using LitvaKebabs.Components;
using System.Linq;
using Microsoft.AspNetCore.Components;
using System.IO;

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
            var config = new CsvConfiguration(System.Globalization.CultureInfo.InvariantCulture)
            {
                HasHeaderRecord = true,
            };
            using var reader = new StreamReader("\"C:\\Users\\AidanFell\\apprenticeship-2023-sd-hci\\LitvaKebabs\\KebabMenu.csv\"");
            using var csv = new CsvReader(reader, config);
            var records = csv.GetRecords<KebabMenuModel>().ToList();
            Dictionary<string, decimal> menuItems = records.ToDictionary(keySelector: records => records.Item, elementSelector: records => records.Price);
            return menuItems;
        }
        private Dictionary<string, decimal> _sauceItems()
        {
            var config = new CsvConfiguration(System.Globalization.CultureInfo.InvariantCulture)
            {
                HasHeaderRecord = true,
            };
            using var reader = new StreamReader(Environment.CurrentDirectory + "\"C:\\Users\\AidanFell\\apprenticeship-2023-sd-hci\\LitvaKebabs\\SauceMenu.csv\"");
            using var csv = new CsvReader(reader, config);
            var records = csv.GetRecords<SauceMenuModel>().ToList();
            Dictionary<string, decimal> sauceItems = records.ToDictionary(keySelector: records => records.Sauce, elementSelector: records => records.Price);
            return sauceItems;
        }
    }
}
