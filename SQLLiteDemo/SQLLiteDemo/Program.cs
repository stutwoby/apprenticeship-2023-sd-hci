using SQLLiteDemo;
using SQLLiteDemo.Models;


// Seed
var database = new DatabaseService();
var menuItemTable = database.GetMenuItemTable();
var orderTable = database.GetOrderTable();

if (menuItemTable == null) throw new Exception();

var hasItems = menuItemTable.Query().ToList().Any();
if (!hasItems)
{
    menuItemTable.InsertBulk(new List<MenuItem> 
    { new MenuItem
        {
            Name = "Kebab",
            Price = 7.99M
        } 
    });
}

var menuItem = menuItemTable.Query().First();

orderTable.Insert(new Order
{
    Name = "Customer Order 1",
    UserID = 1,
    MenuItems = new List<MenuItem> { menuItem }
});

Console.WriteLine(menuItem.Name);
Console.WriteLine(menuItem.Price);

Console.ReadLine();