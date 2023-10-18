using CsvHelper;
using CsvHelper.Configuration;
using LitvaKebabs.Models;
using System.Linq;

namespace LitvaKebabs.Pages
{
    public partial class OrderMenuRoot
    {
        private Dictionary<string, decimal> MenuItems()
        {
            var config = new CsvConfiguration(System.Globalization.CultureInfo.InvariantCulture)
            {
                HasHeaderRecord = true,
            };
            using var reader = new StreamReader("KebabMenu.csv");
            using var csv = new CsvReader(reader, config);
            var records = csv.GetRecords<KebabMenuModel>().ToList();
            Dictionary<string, decimal> menuItems = records.ToDictionary(keySelector: records => records.Item, elementSelector: records => records.Price);
            return menuItems;
        }
        private Dictionary<string, decimal> SauceItems()
        {
            var config = new CsvConfiguration(System.Globalization.CultureInfo.InvariantCulture)
            {
                HasHeaderRecord = true,
            };
            using var reader = new StreamReader("SauceMenu.csv");
            using var csv = new CsvReader(reader, config);
            var records = csv.GetRecords<SauceMenuModel>().ToList();
            Dictionary<string, decimal> sauceItems = records.ToDictionary(keySelector: records => records.Item, elementSelector: records => records.Price);
            return sauceItems;
        }
    }
}
