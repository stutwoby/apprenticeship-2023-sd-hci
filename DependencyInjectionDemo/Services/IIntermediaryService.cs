namespace DependencyInjectionDemo.Services
{
    public interface IIntermediaryService
    {
        public ILogger Logger { get; }

        public IMenuService MenuService { get; }
    }
}
