using Microsoft.EntityFrameworkCore;
using StudydotNETwithPostSQL.Models;
namespace StudydotNETwithPostSQL.Context
{
    public class DbContextApplication :DbContext
    {
        public DbContextApplication(DbContextOptions<DbContextApplication> options) : base(options)
        {

        }
        public DbSet<Student>Students { get; set; }
    }
}
