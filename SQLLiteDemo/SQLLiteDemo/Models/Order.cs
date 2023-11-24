namespace SQLLiteDemo.Models
{
    public class Order
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public int UserID { get; set; }
        public List<MenuItem> MenuItems { get; set; }
    }
}
