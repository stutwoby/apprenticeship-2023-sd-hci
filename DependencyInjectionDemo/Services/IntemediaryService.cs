namespace DependencyInjectionDemo.Services
{
    public class IntemediaryService : IIntermediaryService
    {
        public IntemediaryService(ILogger logger, IMenuService menuService) { 
            Logger = logger;
            MenuService = menuService;
        }
        public ILogger Logger { get; }

        public IMenuService MenuService { get; }
    }
}
