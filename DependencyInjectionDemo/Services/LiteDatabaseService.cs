using LiteDB;

namespace DependencyInjectionDemo.Services
{
    public class LiteDatabaseService : IDatabaseService
    {
        private readonly LiteDatabase _liteDatabase;
        public LiteDatabaseService() {
            _liteDatabase =  new LiteDatabase(@"./mydatabase");
        }

        public ILiteDatabase GetDatabase()
        {
            return _liteDatabase;
        }
    }
}
