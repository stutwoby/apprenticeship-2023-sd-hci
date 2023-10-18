using CsvHelper;
using CsvHelper.Configuration;
using LitvaKebabs.Models;
using System.Linq;

namespace LitvaKebabs.Pages
{
    public partial class OrderMenuRoot
    {
        private Dictionary<string, decimal> GetMenuItems()
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
    }
}
