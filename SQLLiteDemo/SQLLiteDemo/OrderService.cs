using LiteDB;
using SQLLiteDemo.Models;

namespace SQLLiteDemo
{
    public class OrderService
    {
        private DatabaseService _databaseService;
        private ILiteCollection<Order> _orderTable;
        public OrderService() {
            _databaseService = new DatabaseService();
            _orderTable = _databaseService.GetOrderTable();
        }

        public Order GetOrder(int id)
        {
            return _orderTable.Query().Where(x => x.Id == id).First();
        }

        public void UpdateOrder(Order order)
        {
            _orderTable.Update(order);
        }

        public void DeleteOrder(int id)
        {
            _orderTable.Delete(id);
        }
    }
}
