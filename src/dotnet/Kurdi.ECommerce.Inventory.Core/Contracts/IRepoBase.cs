using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using System.Threading.Tasks;

namespace Kurdi.ECommerce.Inventory.Core.Contracts
{
    public interface IRepoBase<T>
    {
        IQueryable<T> FindAll(int pageSize, int pageNumber);
        IQueryable<T> FindAll();
        IQueryable<T> Find(Expression<Func<T, bool>> expression, int pageSize, int pageNumber);
        IQueryable<T> Find(Expression<Func<T, bool>> expression);

        Task Create(T entity);
        Task BulkCreate(List<T> entities);
        void Update(T entity);
        void Delete(T entity);
        int Count();
        Task SaveChanges();

    }
}
