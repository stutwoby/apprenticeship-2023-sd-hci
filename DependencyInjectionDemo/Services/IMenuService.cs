using DependencyInjectionDemo.Models;

namespace DependencyInjectionDemo.Services
{
    public interface IMenuService
    {
        public List<MenuItem> GetAllMenuItems();
    }
}
