using Kurdi.ECommerce.Inventory.Core.Contracts;
using Kurdi.ECommerce.Inventory.Infrastructure.Data;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Threading.Tasks;

namespace Kurdi.ECommerce.Inventory.Infrastructure.DataAccess
{
    public class RepoBase<T> : IRepoBase<T> where T : class
    {
        private readonly AppDbContext _db;

        public RepoBase(AppDbContext db)
        {
            this._db = db;
        }
        public async Task Create(T entity)
        {
            await _db.Set<T>().AddAsync(entity);
        }

        public async Task BulkCreate(List<T> entities)
        {
            await _db.Set<T>().AddRangeAsync(entities);
        }

        public void Delete(T entity)
        {
            this._db.Set<T>().Remove(entity);
        }

        public IQueryable<T> FindAll(int pageSize, int pageNumber)
        {
            return this._db.Set<T>()
                .AsNoTracking()
                .Skip((pageNumber - 1) * pageSize)
                .Take(pageSize);
        }

        public IQueryable<T> FindAll()
        {
            return this._db.Set<T>()
                .AsNoTracking();
        }

        public int Count()
        {
            return this._db.Set<T>()
                .Count();
        }


        public IQueryable<T> Find(Expression<Func<T, bool>> expression, int pageSize, int pageNumber)
        {
            return this._db.Set<T>().Where(expression)
                .AsNoTracking()
                .Skip((pageNumber - 1) * pageSize)
                .Take(pageSize);
        }
        public IQueryable<T> Find(Expression<Func<T, bool>> expression)
        {
            return this._db.Set<T>().Where(expression).AsNoTracking();
        }



        public void Update(T entity)
        {
            this._db.Set<T>().Update(entity);
        }
        public async Task SaveChanges()
        {
            await _db.SaveChangesAsync();
        }
    }
}
