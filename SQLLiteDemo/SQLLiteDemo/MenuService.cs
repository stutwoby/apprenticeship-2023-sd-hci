using LiteDB;
using SQLLiteDemo.Models;

namespace SQLLiteDemo
{
    public class MenuService
    {
        private DatabaseService _databaseService;
        private ILiteCollection<MenuItem> _MenuItemTable;
        public MenuService() {
            _databaseService = new DatabaseService();
            _MenuItemTable = _databaseService.GetMenuItemTable();
        }

        public MenuItem GetMenuItem(int id)
        {
            return _MenuItemTable.Query().Where(x => x.Id == id).First();
        }

        public void UpdateMenuItem(MenuItem MenuItem)
        {
            _MenuItemTable.Update(MenuItem);
        }

        public void DeleteMenuItem(int id)
        {
            _MenuItemTable.Delete(id);
        }
    }
}
