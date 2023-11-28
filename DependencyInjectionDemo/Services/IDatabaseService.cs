using LiteDB;

namespace DependencyInjectionDemo.Services
{
    public interface IDatabaseService
    {
        public ILiteDatabase GetDatabase();
    }
}
