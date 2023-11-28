using DependencyInjectionDemo.Models;
using DependencyInjectionDemo.Services;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace DependencyInjectionDemo.Pages
{
    public class IndexModel : PageModel
    {
        private readonly ILogger<IndexModel> _logger;
        private readonly IMenuService _menuService;
        private readonly IIntermediaryService _intermediaryService;

        public IndexModel(IIntermediaryService intermediaryService)
        {
            _intermediaryService = intermediaryService;
        }

        public List<MenuItem> GetAllMenuitems()
        {
            return _intermediaryService.MenuService.GetAllMenuItems();
        }
    }
}