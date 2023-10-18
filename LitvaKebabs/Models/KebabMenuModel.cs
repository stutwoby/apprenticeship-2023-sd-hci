using CsvHelper.Configuration.Attributes;

namespace LitvaKebabs.Models
{
    public class KebabMenuModel
    {
        [Name("Item")]
        public string Item { get; set; } = string.Empty;

        [Name("Price")]
        public decimal Price { get; set; }
    }
}
