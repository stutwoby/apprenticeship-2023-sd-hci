using DependencyInjectionDemo.Models;
using LiteDB;

namespace DependencyInjectionDemo.Services
{
    public class MenuService : IMenuService
    {
        private readonly ILiteDatabase _liteDatabase;
        private readonly ILiteCollection<MenuItem> _menuItems;

        public MenuService(IDatabaseService database)
        {
            _liteDatabase = database.GetDatabase();
            _menuItems = _liteDatabase.GetCollection<MenuItem>();

            // Seed
            if(!_menuItems.Query().ToList().Any())
            {
                _menuItems.InsertBulk(new MenuItem[]
                {
                    new MenuItem()
                    {
                        Name = "Kebab",
                        Price = 7.99M
                    },
                    new MenuItem()
                    {
                        Name = "Mint Sauce",
                        Price = 0.5M
                    },
                    new MenuItem()
                    {
                        Name = "Coke",
                        Price = 1.99M
                    },
                    new MenuItem()
                    {
                        Name = "Chips",
                        Price = 2.99M
                    },
                });
            }
        }

        public List<MenuItem> GetAllMenuItems()
        {
            return _menuItems.Query().ToList();
        }
    }
}
