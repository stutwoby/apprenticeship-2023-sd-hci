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
            CsvConfiguration config = new CsvConfiguration(System.Globalization.CultureInfo.InvariantCulture)
            {
                HasHeaderRecord = true,
            };
            try
            {
                FileInfo fi = new(@"\KebabMenu.csv");
                using FileStream fs = fi.Open(FileMode.Open, FileAccess.Read, FileShare.Read);
                using StreamReader reader = new StreamReader(fs);
                using CsvReader csv = new CsvReader(reader, config);
                List<KebabMenuModel> records = csv.GetRecords<KebabMenuModel>().ToList();
                Dictionary<string, decimal> menuItems = records.ToDictionary(keySelector: records => records.Item, elementSelector: records => records.Price);
                return menuItems;
            }
            catch (FileNotFoundException f)
            {
                throw;
            }

        }
        private Dictionary<string, decimal> _sauceItems()
        {
            CsvConfiguration config = new CsvConfiguration(System.Globalization.CultureInfo.InvariantCulture)
            {
                HasHeaderRecord = true,
            };
            string saucePath = "SauceMenu.csv";
            //var testSaucePath = System.IO.Directory.GetFiles("/");
            using StreamReader reader = new StreamReader(saucePath);
            using CsvReader csv = new CsvReader(reader, config);
            List<SauceMenuModel> records = csv.GetRecords<SauceMenuModel>().ToList();
            Dictionary<string, decimal> sauceItems = records.ToDictionary(keySelector: records => records.Sauce, elementSelector: records => records.Price);
            return sauceItems;
        }
    }
}
